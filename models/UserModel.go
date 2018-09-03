package models

import (
		"time"
	"upper.io/db.v3"
	"github.com/jtsalva/estore/auth"
	"errors"
)

type users struct {}

type User struct {
	Id int64 `db:"UserId" json:"id"`
	Name string `db:"Name" json:"name"`
	Email string `db:"Email" json:"email"`
	Password string `db:"Password" json:"-"` // Omit from json - otherwise can be accessed via api
	DataJoined time.Time `db:"DateJoined" json:"dateJoined"`
	RoleId int64 `db:"Role" json:"roleId"`
}

var Users users

func (u *users) ALl() (*[]User, error) {
	users, err := all(User{})
	return users.(*[]User), err
}

func (u *users) GetById(id int64) (*User, error) {
	user, err := getById(User{}, id)
	return user.(*User), err
}

func (u *users) GetByEmail(email string) (*User, error) {
	sess := newSession()
	defer sess.Close()

	var user User

	err := sess.SelectFrom(UsersTable).Where(db.Cond{"Email": email}).One(&user)
	return &user, err
}

func (u *users) Insert(user User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.DataJoined = time.Now()

	return insert(user)
}

func (u *users) RemoveById(id int64) error {
	return removeById(User{}, id)
}

func (u *User) Role() (Role, error) {
	sess := newSession()
	defer sess.Close()

	var role Role

	err := sess.SelectFrom(RolesTable).Where(db.Cond{"RoleId": u.RoleId}).One(&role)
	return role, err
}

func (u *User) Authenticate() (bool, error) {
	var storedUser *User
	var err error

	if u.Password == "" {
		return false, errors.New("password is missing or empty")
	}

	if u.Id != 0 {
		storedUser, err = Users.GetById(u.Id)
	} else if u.Email != "" {
		storedUser, err = Users.GetByEmail(u.Email)
	} else {
		return false, errors.New("missing parameters for user authentication")
	}

	if err != nil {
		return false, err
	}

	return auth.PasswordMatchesHash(u.Password, storedUser.Password), nil
}

func (u User) Update() error {
	return update(u)
}