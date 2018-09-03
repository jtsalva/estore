package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	"log"
	"github.com/jtsalva/estore/models"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./content/static"))))

	http.ListenAndServe(":8000", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index, err := template.ParseFiles("./content/templates/index.html")
	if err != nil {
		log.Println(err.Error())
		return
	}

	items, err := models.Items.All()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for idx, item := range *items {
		if len(item.Description) > 75 {
			(*items)[idx].Description = item.Description[:75]
		}
	}

	categories, err := models.Categories.All()
	if err != nil {
		log.Println(err.Error())
		return
	}

	index.Execute(w, struct {
		Items []models.Item
		Categories []models.Category
	} {
		Items: (*items)[:50],
		Categories: *categories,
	})
}