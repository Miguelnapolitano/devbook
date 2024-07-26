package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/models/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Creates an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("POST"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	userID, err := repo.Creates(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	user.ID = userID
	responses.JSON(w, http.StatusCreated, user)
}

// List all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	userList, err := repo.List(nameOrNick)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userList)

}

// Retrieves an user
func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	user, err := repo.RetrieveUser(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)

}

// Updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("PUT"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	if err = repo.UpdateUser(userID, user); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// Deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	if err = repo.DeleteUser(userID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
