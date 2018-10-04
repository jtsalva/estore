package request

import (
	"github.com/jtsalva/estore/models"
	"github.com/jinzhu/copier"
)

func tagModel(req interface{}) *models.Tag {
	var tag models.Tag

	copier.Copy(&tag, req)
	return &tag
}

type GetTagRequest struct {
	Id int64 `json:"id"`
}

type CreateTagRequest struct {
	Name string `json:"name" required:"true"`
}

func (c *CreateTagRequest) Model() *models.Tag {
	return tagModel(c)
}

type UpdateTagRequest struct {
	Id int64 `json:"id" required:"true"`
	Name string `json:"id"`
}

func (u *UpdateTagRequest) Model() *models.Tag {
	return tagModel(u)
}

type DeleteTagRequest struct {
	Id int64 `json:"id" required:"true"`
}

