package controllers

import (
	"encoding/json"
	"heartvoice/src/database"
	"heartvoice/src/models"
	"heartvoice/src/repositories"
	"heartvoice/src/response"
	"io"
	"net/http"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	requestBody, bodyError := io.ReadAll(r.Body)

	if bodyError != nil {
		response.Error(w, http.StatusUnprocessableEntity, bodyError)
		return
	}

	var user models.User

	if jsonError := json.Unmarshal(requestBody, &user); jsonError != nil {
		response.Error(w, http.StatusBadRequest, jsonError)
		return
	}

	if prepareError := user.Prepare("signup"); prepareError != nil {
		response.Error(w, http.StatusBadRequest, prepareError)
		return
	}

	db, dbError := database.Connect()

	if dbError != nil {
		response.Error(w, http.StatusInternalServerError, dbError)
		return
	}

	defer db.Close()

	repository := repositories.UserRepository(db)

	user.ID, dbError = repository.Create(user)
	if dbError != nil {
		response.Error(w, http.StatusInternalServerError, dbError)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}
