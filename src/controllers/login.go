package controllers

import (
	"encoding/json"
	"heartvoice/src/authentication"
	"heartvoice/src/database"
	"heartvoice/src/models"
	"heartvoice/src/repositories"
	"heartvoice/src/response"
	"heartvoice/src/security"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, bodyError := io.ReadAll(r.Body)

	if bodyError != nil {
		response.Error(w, http.StatusUnprocessableEntity, bodyError)
		return
	}

	var user models.User
	if bodyError = json.Unmarshal(requestBody, &user); bodyError != nil {
		response.Error(w, http.StatusBadRequest, bodyError)
		return
	}

	db, dbError := database.Connect()
	if dbError != nil {
		response.Error(w, http.StatusInternalServerError, dbError)
		return
	}

	defer db.Close()

	repository := repositories.UserRepository(db)

	databaseUser, dbError := repository.FindBy("email", user.Email)

	if dbError != nil {
		response.Error(w, http.StatusInternalServerError, dbError)
		return
	}

	validatePassswordError := security.ValidatePassword(user.Password, databaseUser.Password)
	if validatePassswordError != nil {
		response.Error(w, http.StatusUnauthorized, validatePassswordError)
		return
	}

	token, authError := authentication.CreateToken(databaseUser.ID)
	if authError != nil {
		response.Error(w, http.StatusInternalServerError, authError)
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
