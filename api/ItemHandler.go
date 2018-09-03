package main

import (
	"net/http"
	"github.com/jtsalva/estore/models"
	"github.com/jtsalva/estore/api/request"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := models.Items.All()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	var getRequest request.GetItemRequest

	if err := validateRequest(r, &getRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	item, err := models.Items.GetById(getRequest.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var createRequest request.CreateItemRequest

	if err := validateRequest(r, &createRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Items.Insert(*createRequest.Model()); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updateRequest request.UpdateItemRequest

	if err := validateRequest(r, &updateRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := updateRequest.Model().Update(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	var deleteRequest request.DeleteItemRequest

	if err := validateRequest(r, &deleteRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Items.RemoveById(deleteRequest.Id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}