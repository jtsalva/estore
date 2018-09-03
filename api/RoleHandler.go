package main

import (
	"net/http"
	"github.com/jtsalva/estore/api/request"
	"github.com/jtsalva/estore/models"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := models.Roles.All()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, roles)
}

func GetRole(w http.ResponseWriter, r *http.Request) {
	var getRequest request.GetRoleRequest

	if err := validateRequest(r, &getRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	role, err := models.Roles.GetById(getRequest.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, role)
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	var createRequest request.CreateRoleRequest

	if err := validateRequest(r, &createRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Roles.Insert(*createRequest.Model()); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	var updateRequest request.UpdateRoleRequest

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

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	var deleteRequest request.DeleteRoleRequest

	if err := validateRequest(r, &deleteRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Roles.RemoveById(deleteRequest.Id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}