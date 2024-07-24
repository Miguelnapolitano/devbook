package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/models/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

//Creates an user
func CreateUser(w http.ResponseWriter, r *http.Request){
	requestBody, err := io.ReadAll(r.Body)	
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err  != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
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

//List all users
func ListUsers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Listing all user"))

}

//Retrieves an user
func RetrieveUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Retrieving an user"))

}

//Updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Updating an user"))

}

//Deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deleting an user"))

}