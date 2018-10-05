package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jtsalva/estore/models"
)

const (
	Domain string = "shop.jtsalva.space"

	Certificate string = "/etc/letsencrypt/live/shop.jtsalva.space/cert.pem"
	PrivateKey  string = "/etc/letsencrypt/live/shop.jtsalva.space/privkey.pem"
)

func main() {
	r := mux.NewRouter().Host(Domain).Subrouter()

	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./content/static"))))

	// Redirect port 80 traffic to 443
	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	http.ListenAndServeTLS(":443", Certificate, PrivateKey, r)
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+Domain+r.RequestURI, http.StatusMovedPermanently)
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
		Items      []models.Item
		Categories []models.Category
	}{
		Items:      (*items)[:75],
		Categories: *categories,
	})
}
