package controllers

import (
	"encoding/json"
	"heartvoice/src/database"
	"heartvoice/src/models"
	"heartvoice/src/repositories"
	"heartvoice/src/response"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateGuild(w http.ResponseWriter, r *http.Request) {
	requestBody, bodyError := io.ReadAll(r.Body)

	if bodyError != nil {
		response.Error(w, http.StatusUnprocessableEntity, bodyError)
		return
	}

	var guild models.Guild

	if jsonError := json.Unmarshal(requestBody, &guild); jsonError != nil {
		response.Error(w, http.StatusBadRequest, jsonError)
		return
	}

	if prepareError := guild.Prepare(); prepareError != nil {
		response.Error(w, http.StatusBadRequest, prepareError)
		return
	}

	db, dbError := database.Connect()
	if dbError != nil {
		response.Error(w, http.StatusInternalServerError, dbError)
		return
	}

	defer db.Close()

	repository := repositories.GuildRepository(db)

	guild.ID, dbError = repository.Create(guild)
	if dbError != nil {
		response.Error(w, http.StatusInternalServerError, dbError)
		return
	}

	response.JSON(w, http.StatusCreated, guild)
}

func FindByGuild(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	guildId, convError := strconv.ParseUint(parameters["id"], 10, 64)
	if convError != nil {
		response.Error(w, http.StatusBadRequest, convError)
		return
	}

	db, dbError := database.Connect()
	if dbError != nil {
		response.Error(w, http.StatusInternalServerError, dbError)
		return
	}

	defer db.Close()

	repository := repositories.GuildRepository(db)

	guild, dbError := repository.FindBy("id", guildId)

	if dbError != nil {
		response.Error(w, http.StatusInternalServerError, dbError)
		return
	}

	response.JSON(w, http.StatusOK, guild)
}

func FindAllGuilds(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("name"))

	limit, limitError := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
	page, pageError := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)

	if limitError != nil || pageError != nil {
		limit = 10
		page = 0
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.GuildRepository(db)

	pagination := repositories.PaginationParams{
		Page:  page,
		Limit: limit,
	}

	guilds, rpError := repository.FindAll(pagination, name)
	if rpError != nil {
		response.Error(w, http.StatusInternalServerError, rpError)
		return
	}

	response.JSON(w, http.StatusOK, guilds)
}
