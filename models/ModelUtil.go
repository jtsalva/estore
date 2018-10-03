package models

import (
	"errors"
	"fmt"
	"reflect"

	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var settings = mysql.ConnectionURL{
	Database: "store",
	Host:     "138.68.175.121",
	User:     "root",
	Password: "password",
	// Very secure way to access
	// the database I promise
}

// Tables
const (
	TagsTable       string = "Tags"
	CategoriesTable string = "Categories"
	RolesTable      string = "Roles"
	UsersTable      string = "Users"
	ItemsTable      string = "Items"
	ItemsTagsTable  string = "ItemTags"
)

// TODO: Handle no connection to database
func newSession() (sqlbuilder.Database, error) {
	sess, err := mysql.Open(settings)
	if err != nil {
		return sess, err
	}

	return sess, nil
}

// Concatenate Id onto the name of model
func idString(modelType interface{}) string {
	return fmt.Sprintf("%sId", reflect.TypeOf(modelType).Name())
}

func isValueEmpty(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.String() == ""
	case reflect.Int64:
		return val.Int() == 0
	case reflect.Float64:
		return val.Float() == 0
	default:
		// Default case indicates
		// the type is unused within
		// database schema
		return true
	}
}

func tableName(modelType interface{}) (string, error) {
	switch modelType.(type) {
	case Category:
		return CategoriesTable, nil
	case Item:
		return ItemsTable, nil
	case Role:
		return RolesTable, nil
	case Tag:
		return TagsTable, nil
	case User:
		return UsersTable, nil
	default:
		return "", errors.New("tableName(): invalid modelType")
	}
}

func LinkItemWithTag(itemId int64, tagId int64) error {
	sess, err := newSession()
	if err != nil {
		return err
	}
	defer sess.Close()

	_, err = sess.InsertInto(ItemsTagsTable).Values(struct {
		ItemId int64 `db:"ItemId"`
		TagId  int64 `db:"TagId"`
	}{
		itemId,
		tagId,
	}).Exec()
	return err
}

func UnlinkItemWithTag(itemId int64, tagId int64) error {
	sess, err := newSession()
	if err != nil {
		return err
	}
	defer sess.Close()

	_, err = sess.DeleteFrom(ItemsTagsTable).Where(
		db.And(
			db.Cond{"ItemId": itemId},
			db.Cond{"TagId": tagId},
		)).Exec()
	return err
}

func all(modelType interface{}) (interface{}, error) {
	models := reflect.New(reflect.SliceOf(reflect.TypeOf(modelType))).Interface()

	sess, err := newSession()
	if err != nil {
		return models, err
	}
	defer sess.Close()

	// Create new slice of model type || reflect.New returns address
	table, err := tableName(modelType)
	if err != nil {
		return models, err
	}

	err = sess.SelectFrom(table).All(models)
	return models, err
}

func getById(modelType interface{}, id int64) (interface{}, error) {
	sess, err := newSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	model := reflect.New(reflect.TypeOf(modelType)).Interface()
	table, err := tableName(modelType)
	if err != nil {
		return model, nil
	}

	err = sess.SelectFrom(table).Where(db.Cond{idString(modelType): id}).One(model)
	return model, err
}

// Only used by specific models
func getByName(modelType interface{}, name string) (interface{}, error) {
	sess, err := newSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	model := reflect.New(reflect.TypeOf(modelType)).Interface()
	table, err := tableName(modelType)
	if err != nil {
		return model, nil
	}

	err = sess.SelectFrom(table).Where(db.Cond{"Name": name}).One(model)
	return model, err
}

func insert(model interface{}) error {
	sess, err := newSession()
	if err != nil {
		return err
	}
	defer sess.Close()

	table, err := tableName(model)
	if err != nil {
		return err
	}

	_, err = sess.InsertInto(table).Values(model).Exec()
	return err
}

func update(model interface{}) error {
	sess, err := newSession()
	if err != nil {
		return err
	}
	defer sess.Close()

	table, err := tableName(model)
	if err != nil {
		return err
	}

	newData := make(map[string]interface{})

	// Filter empty struct values
	// Only new fields will be updated
	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i)
		columnName := v.Type().Field(i).Tag.Get("db")
		if !isValueEmpty(val) {
			newData[columnName] = val.Interface()
		}
	}

	// If the id exists from the model
	if id, ok := newData[idString(model)]; ok {
		_, err = sess.Update(table).Set(newData).Where(db.Cond{idString(model): id}).Exec()
	} else {
		return err
	}

	return err
}

func removeById(modelType interface{}, id int64) error {
	sess, err := newSession()
	if err != nil {
		return err
	}
	defer sess.Close()

	table, err := tableName(modelType)
	if err != nil {
		return err
	}

	_, err = sess.DeleteFrom(table).Where(db.Cond{idString(modelType): id}).Exec()
	return err
}

func removeByName(modelType interface{}, name string) error {
	sess, err := newSession()
	if err != nil {
		return err
	}
	defer sess.Close()

	table, err := tableName(modelType)
	if err != nil {
		return err
	}

	_, err = sess.DeleteFrom(table).Where(db.Cond{"Name": name}).Exec()
	return err
}
