package main

import (
	"github.com/jtsalva/estore/models"
	"log"
	"encoding/json"
)

func main() {
	all, err := models.Categories.All()
	if err != nil {
		log.Println(err.Error())
	}

	payload, err := json.Marshal(all)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(string(payload))
}