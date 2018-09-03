package models

import (
	"log"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
		"errors"
	"reflect"
	"upper.io/db.v3"
	"fmt"
)

var settings = mysql.ConnectionURL{
	Database: "store",
	Host: "localhost",
	User: "root",
	// Optional Password
}

// Tables
const (
	TagsTable       string = "tags"
	CategoriesTable string = "categories"
	RolesTable      string = "roles"
	UsersTable      string = "users"
	ItemsTable      string = "items"
	ItemsTagsTable  string = "itemtags"
)

// TODO: Handle no connection to database
func newSession() sqlbuilder.Database {
	sess, err := mysql.Open(settings)
	if err != nil {
		log.Println("newSession(): %q\n", err)
	}

	return sess
}

// Concatenate Id onto the name of model
func idString(modelType interface{}) string {
	return fmt.Sprintf("%sId", reflect.TypeOf(modelType).Name())
}

// Changed this so check functionality please
func isValueEmpty(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.String() == ""
	case reflect.Int64:
		return val.Int() == 0
	case reflect.Float64:
		return val.Float() == 0
	default:
		// Type unused in database so return true
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
		return "", errors.New("can't find model type")
	}
}

func LinkItemWithTag(itemId int64, tagId int64) error {
	sess := newSession()
	defer sess.Close()

	_, err := sess.InsertInto(ItemsTagsTable).Values(struct{
		ItemId int64 `db:"ItemId"`
		TagId int64 `db:"TagId"`
	}{
		itemId,
		tagId,
	}).Exec()
	return err
}

func UnlinkItemWithTag(itemId int64, tagId int64) error {
	sess := newSession()
	defer sess.Close()

	_, err := sess.DeleteFrom(ItemsTagsTable).Where(
		db.And(
			db.Cond{"ItemId": itemId},
			db.Cond{"TagId": tagId},
			)).Exec()
	return err
}

func all(modelType interface{}) (interface{}, error) {
	sess := newSession()
	defer sess.Close()

	// CreateOne new slice of model type || reflect.New returns address already
	models := reflect.New(reflect.SliceOf(reflect.TypeOf(modelType))).Interface()
	table, err := tableName(modelType)
	if err != nil {
		return models, err
	}

	err = sess.SelectFrom(table).All(models)
	return models, err
}

func getById(modelType interface{}, id int64) (interface{}, error) {
	sess := newSession()
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
	sess := newSession()
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
	sess := newSession()
	defer sess.Close()

	table, err := tableName(model)
	if err != nil {
		return err
	}

	_, err = sess.InsertInto(table).Values(model).Exec()
	return err
}

func update(model interface{}) error {
	sess := newSession()
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
		return errors.New("missing id in update model")
	}

	return err
}

func removeById(modelType interface{}, id int64) error {
	sess := newSession()
	defer sess.Close()

	table, err := tableName(modelType)
	if err != nil {
		return err
	}

	_, err = sess.DeleteFrom(table).Where(db.Cond{idString(modelType): id}).Exec()
	return err
}

func removeByName(modelType interface{}, name string) error {
	sess := newSession()
	defer sess.Close()

	table, err := tableName(modelType)
	if err != nil {
		return err
	}

	_, err = sess.DeleteFrom(table).Where(db.Cond{"Name": name}).Exec()
	return err
}