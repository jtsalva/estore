package request

import (
	"github.com/jtsalva/estore/models"
	"github.com/jinzhu/copier"
)

func categoryModel(req interface{}) *models.Category {
	var category models.Category

	copier.Copy(&category, req)
	return &category
}

type GetCategoryRequest struct {
	Id int64 `json:"id"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" required:"true"`
}

func (c *CreateCategoryRequest) Model() *models.Category {
	return categoryModel(c)
}

type UpdateCategoryRequest struct {
	Id int64 `json:"id" required:"true"`
	Name string `json:"name"`
}

func (u *UpdateCategoryRequest) Model() *models.Category {
	return categoryModel(u)
}

type DeleteCategoryRequest struct {
	Id int64 `json:"id" required:"true"`
}
