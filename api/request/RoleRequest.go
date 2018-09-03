package request

import (
	"github.com/jtsalva/estore/models"
	"github.com/jinzhu/copier"
)

func roleModel(req interface{}) *models.Role {
	var role models.Role

	copier.Copy(&role, req)
	return &role
}

type GetRoleRequest struct {
	Id int64 `json:"id"`
}

type CreateRoleRequest struct {
	Name string `json:"name" required:"true"`
}

func (c *CreateRoleRequest) Model() *models.Role {
	return roleModel(c)
}

type UpdateRoleRequest struct {
	Id int64 `json:"id" required:"true"`
	Name int64 `json:"name"`
}

func (u *UpdateRoleRequest) Model() *models.Role {
	return roleModel(u)
}

type DeleteRoleRequest struct {
	Id int64 `json:"id" required:"true"`
}
