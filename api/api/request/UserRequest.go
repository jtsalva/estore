package request

import (
	"github.com/jtsalva/estore/models"
	"github.com/jinzhu/copier"
)

func userModel(req interface{}) *models.User {
	var user models.User

	copier.Copy(&user, req)
	return &user
}

type GetUserRequest struct {
	Id int64 `json:"id"`
}

type CreateUserRequest struct {
	Name string `json:"name" required:"true"`
	Email string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
	RoleId int64 `json:"roleId"`
}

func (c *CreateUserRequest) Model() *models.User {
	return userModel(c)
}

type DeleteUserRequest struct {
	Id int64 `json:"id" required:"true"`
}

type UpdateUserRequest struct {
	Id int64 `json:"id" required:"true"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role int64 `json:"role"`
}

func (u *UpdateUserRequest) Model() *models.User {
	return userModel(u)
}

type AuthenticateUserRequest struct {
	Email string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

func (a *AuthenticateUserRequest) Model() *models.User {
	return userModel(a)
}