package models

type categories struct{}

type Category struct {
	Id   int64  `db:"CategoryId"`
	Name string `db:"Name"`
}

var Categories categories

func (c *categories) All() (*[]Category, error) {
	categories, err := all(Category{})
	return categories.(*[]Category), err
}

func (c *categories) GetById(id int64) (*Category, error) {
	category, err := getById(Category{}, id)
	return category.(*Category), err
}

func (c *categories) GetByName(name string) (*Category, error) {
	category, err := getByName(Category{}, name)
	return category.(*Category), err
}

func (c *categories) Insert(category Category) error {
	return insert(category)
}

func (c *categories) RemoveById(id int64) error {
	return removeById(Category{}, id)
}

func (c *categories) RemoveByName(name string) error {
	return removeByName(Category{}, name)
}

func (c Category) Update() error {
	return update(c)
}
