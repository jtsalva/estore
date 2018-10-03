package models

type tags struct{}

type Tag struct {
	Id   int64  `db:"TagId" json:"id"`
	Name string `db:"Name" json:"name"`
}

var Tags tags

func (t *tags) All() (*[]Tag, error) {
	tags, err := all(Tag{})
	return tags.(*[]Tag), err
}

func (t *tags) GetById(id int64) (*Tag, error) {
	tag, err := getById(Tag{}, id)
	return tag.(*Tag), err
}

func (t *tags) GetByName(name string) (*Tag, error) {
	tag, err := getByName(Tag{}, name)
	return tag.(*Tag), err
}

func (t *tags) Insert(tag Tag) error {
	return insert(tag)
}

func (t *tags) RemoveById(id int64) error {
	return removeById(Tag{}, id)
}

func (t *tags) RemoveByName(name string) error {
	return removeByName(Tag{}, name)
}

func (t Tag) Update() error {
	return update(t)
}
