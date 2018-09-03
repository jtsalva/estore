package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/jtsalva/estore/api/request"
	"github.com/gorilla/mux"
)

// Wrapper to set header response type for all handlers
func handler(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Temporary for testing only
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		h(w, r)
	}
}

func respondWithError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	log.Println(fmt.Sprintf("api: %s", err.Error()))
}

func respondWithData(w http.ResponseWriter, status int, data interface{}) {
	if response, err := json.Marshal(data); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(response)
	}
}

func validateRequest(r *http.Request, req interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if r.Method == http.MethodGet || r.Method == http.MethodDelete ||
		r.Method == http.MethodPatch || r.Method == http.MethodPut {
		id := mux.Vars(r)["id"]

		if id == "" {
			body = []byte("{}")
		} else {
			body = []byte(fmt.Sprintf(`{"id": %s}`, id))
		}
	}

	if err := json.Unmarshal(body, &req); err != nil {
		return err
	}


	if request.IsIncomplete(req) {
		return request.IncompleteRequestError
	}

	return nil
}