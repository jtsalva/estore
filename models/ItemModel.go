package models

import (
	"fmt"

	"upper.io/db.v3"
)

type items struct{}

type Item struct {
	Id          int64  `db:"ItemId" json:"id"`
	Name        string `db:"Name" json:"name"`
	Description string `db:"Description" json:"description"`
	Price       Price  `db:"Price" json:"price"`
	CategoryId  int64  `db:"Category" json:"categoryId"`
}

var Items items

type Price float64

func (p Price) String() string {
	return fmt.Sprintf("%.2f", float64(p))
}

func (i *items) All() (*[]Item, error) {
	items, err := all(Item{})
	return items.(*[]Item), err
}

func (i *items) GetById(id int64) (*Item, error) {
	item, err := getById(Item{}, id)
	return item.(*Item), err
}

func (i *items) Insert(item Item) error {
	return insert(item)
}

func (i *items) RemoveById(id int64) error {
	return removeById(Item{}, id)
}

func (i *Item) Category() (*Category, error) {
	return Categories.GetById(i.CategoryId)
}

func (i *Item) Tags() ([]Tag, error) {
	var tags []Tag

	sess, err := newSession()
	if err != nil {
		return tags, err
	}
	defer sess.Close()

	var itemTagPairs []struct {
		ItemId int64 `db:"ItemId"`
		TagId  int64 `db:"TagId"`
	}

	if err := sess.SelectFrom(ItemsTagsTable).Where(db.Cond{"ItemId": i.Id}).All(&itemTagPairs); err != nil {
		return tags, err
	}

	for _, pair := range itemTagPairs {
		if tag, err := Tags.GetById(pair.TagId); err != nil {
			return tags, err
		} else {
			tags = append(tags, *tag)
		}
	}

	return tags, nil
}

func (i Item) Update() error {
	return update(i)
}
