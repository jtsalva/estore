package main

import (
	"net/http"
	"github.com/jtsalva/estore/api/request"
	"github.com/jtsalva/estore/models"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := models.Categories.All()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, categories)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	var getRequest request.GetCategoryRequest

	if err := validateRequest(r, &getRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	category, err := models.Categories.GetById(getRequest.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, category)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var createRequest request.CreateCategoryRequest

	if err := validateRequest(r, &createRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Categories.Insert(*createRequest.Model()); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var updateRequest request.UpdateCategoryRequest

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

func DeleteCategory(w http.ResponseWriter, r *http.Request){
	var deleteRequest request.DeleteCategoryRequest

	if err := validateRequest(r, &deleteRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Categories.RemoveById(deleteRequest.Id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}