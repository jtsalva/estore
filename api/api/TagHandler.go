package main

import (
	"net/http"
	"github.com/jtsalva/estore/api/request"
	"github.com/jtsalva/estore/models"
)

func GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := models.Tags.All()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, tags)
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	var getRequest request.GetTagRequest

	if err := validateRequest(r, &getRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	tag, err := models.Tags.GetById(getRequest.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, tag)
}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	var createRequest request.CreateTagRequest

	if err := validateRequest(r, &createRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Tags.Insert(*createRequest.Model()); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateTag(w http.ResponseWriter, r *http.Request) {
	var updateRequest request.UpdateTagRequest

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

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	var deleteRequest request.DeleteTagRequest

	if err := validateRequest(r, &deleteRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Tags.RemoveById(deleteRequest.Id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}