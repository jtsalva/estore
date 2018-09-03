package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jtsalva/estore/api/request"
	"fmt"
	"log"
	cors2 "github.com/rs/cors"
)

const (
	//Host string = "localhost"
	Port int32  = 8080
)

var (
	api = mux.NewRouter()

	items request.Path
	users request.Path
	categories request.Path
	tags request.Path
	roles request.Path
)

func main() {
	//api.Host(Host)

	setItemHandlers()
	setUserHandlers()
	setCategoryHandlers()
	setRoleHandlers()
	setTagHandlers()

	cors := cors2.New(cors2.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		//Debug: true,
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), cors.Handler(api)); err != nil {
		log.Fatal(err)
	}
}

func setItemHandlers() {
	items.SetPath("/items/")

	api.Path(items.GetMultiple).HandlerFunc(handler(GetItems)).Methods(http.MethodGet)
	api.Path(items.GetOne).HandlerFunc(handler(GetItem)).Methods(http.MethodGet)
	api.Path(items.CreateOne).HandlerFunc(handler(CreateItem)).Methods(http.MethodPost)
	api.Path(items.UpdateOne).HandlerFunc(handler(UpdateItem)).Methods(http.MethodPatch, http.MethodPut)
	api.Path(items.DeleteOne).HandlerFunc(handler(DeleteItem)).Methods(http.MethodDelete)
}

func setUserHandlers() {
	users.SetPath("/users/")

	api.Path(users.GetMultiple).HandlerFunc(handler(GetUsers)).Methods(http.MethodGet, http.MethodOptions)
	api.Path(users.GetOne).HandlerFunc(handler(GetUser)).Methods(http.MethodGet)
	api.Path(users.CreateOne).HandlerFunc(handler(CreateUser)).Methods(http.MethodPost)
	api.Path(users.UpdateOne).HandlerFunc(handler(UpdateUser)).Methods(http.MethodPatch, http.MethodPut)
	api.Path(users.DeleteOne).HandlerFunc(handler(DeleteUser)).Methods(http.MethodDelete)

	api.Path(users.Endpoint("authenticate")).HandlerFunc(handler(AuthenticateUser)).Methods(http.MethodPost)
}

func setCategoryHandlers() {
	categories.SetPath("/categories/")

	api.Path(categories.GetMultiple).HandlerFunc(handler(GetCategories)).Methods(http.MethodGet)
	api.Path(categories.GetOne).HandlerFunc(handler(GetCategory)).Methods(http.MethodGet)
	api.Path(categories.CreateOne).HandlerFunc(handler(CreateCategory)).Methods(http.MethodPost)
	api.Path(categories.UpdateOne).HandlerFunc(handler(UpdateCategory)).Methods(http.MethodPatch, http.MethodPut)
	api.Path(categories.DeleteOne).HandlerFunc(handler(DeleteCategory)).Methods(http.MethodDelete)
}

func setTagHandlers() {
	tags.SetPath("/tags/")

	api.Path(tags.GetMultiple).HandlerFunc(handler(GetTags)).Methods(http.MethodGet)
	api.Path(tags.GetOne).HandlerFunc(handler(GetTag)).Methods(http.MethodGet)
	api.Path(tags.CreateOne).HandlerFunc(handler(CreateTag)).Methods(http.MethodPost)
	api.Path(tags.UpdateOne).HandlerFunc(handler(UpdateTag)).Methods(http.MethodPatch, http.MethodPut)
	api.Path(tags.DeleteOne).HandlerFunc(handler(DeleteTag)).Methods(http.MethodDelete)
}

func setRoleHandlers() {
	roles.SetPath("/roles/")

	api.Path(roles.GetMultiple).HandlerFunc(handler(GetRoles)).Methods(http.MethodGet)
	api.Path(roles.GetOne).HandlerFunc(handler(GetRole)).Methods(http.MethodGet)
	api.Path(roles.CreateOne).HandlerFunc(handler(CreateRole)).Methods(http.MethodPost)
	api.Path(roles.UpdateOne).HandlerFunc(handler(UpdateRole)).Methods(http.MethodPatch, http.MethodPut)
	api.Path(roles.DeleteOne).HandlerFunc(handler(DeleteRole)).Methods(http.MethodDelete)
}