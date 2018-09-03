package request

import (
	"github.com/jtsalva/estore/models"
	"github.com/jinzhu/copier"
)

func itemModel(req interface{}) *models.Item {
	var item models.Item

	copier.Copy(&item, req)
	return &item
}

type GetItemRequest struct {
	Id int64 `json:"id"`
}

type CreateItemRequest struct {
	Name string `json:"name" required:"true"`
	Description string `json:"description"`
	Price float64 `json:"price" required:"true"`
	CategoryId int64 `json:"categoryId"`
}

func (c *CreateItemRequest) Model() *models.Item {
	return itemModel(c)
}

type DeleteItemRequest struct {
	Id int64 `json:"id" required:"true"`
}

type UpdateItemRequest struct {
	Id int64 `json:"id" required:"true"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CategoryId int64 `json:"categoryId"`
}

func (u *UpdateItemRequest) Model() *models.Item {
	return itemModel(u)
}