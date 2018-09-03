package main

import (
	"net/http"
	"github.com/jtsalva/estore/api/request"
	"github.com/jtsalva/estore/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.Users.ALl()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var getRequest request.GetUserRequest

	if err := validateRequest(r, &getRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := models.Users.GetById(getRequest.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithData(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var createRequest request.CreateUserRequest

	if err := validateRequest(r, &createRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Users.Insert(*createRequest.Model()); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateRequest request.UpdateUserRequest

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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteRequest request.DeleteUserRequest

	if err := validateRequest(r, &deleteRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := models.Users.RemoveById(deleteRequest.Id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var authenticateRequest request.AuthenticateUserRequest

	if err := validateRequest(r, &authenticateRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if authenticated, err := authenticateRequest.Model().Authenticate(); err != nil {
		respondWithError(w, http.StatusUnauthorized, err)
		return
	} else if authenticated {
		if user, err := models.Users.GetByEmail(authenticateRequest.Email); err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
		} else {
			respondWithData(w, http.StatusOK, user)
		}
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}