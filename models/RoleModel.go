package models

type roles struct {}

type Role struct {
	Id int64 `db:"RoleId" json:"id"`
	Name string `db:"Name" json:"name"`
}

var Roles roles

func (r *roles) All() (*[]Role, error) {
	roles, err := all(Role{})
	return roles.(*[]Role), err
}

func (r *roles) GetById(id int64) (*Role, error) {
	role, err := getById(Role{}, id)
	return role.(*Role), err
}

func (r *roles) GetByName(name string) (*Role, error) {
	role, err := getByName(Role{}, name)
	return role.(*Role), err
}

func (r *roles) Insert(role Role) error {
	return insert(role)
}

func (r *roles) RemoveById(id int64) error {
	return removeById(Role{}, id)
}

func (r *roles) RemoveByName(name string) error {
	return removeByName(Role{}, name)
}

func (r Role) Update() error {
	return update(r)
}