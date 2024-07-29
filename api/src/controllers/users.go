package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/models/repositories"
	"api/src/responses"
	"api/src/secure"
	"encoding/json"
	"errors"
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
	
	tokenUserID, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if tokenUserID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("it's not yours user to update"))
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
	
	tokenUserID, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if tokenUserID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("it's not yours user to delete"))
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

// follow and user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Err(w, http.StatusForbidden, errors.New("it's not possible follow you"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	if err = repo.Follow(userID, followerID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Stop follow and user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Err(w, http.StatusForbidden, errors.New("it's not possible unfollow you"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	if err = repo.Unfollow(userID, followerID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Stop follow and user
func ListFollowers(w http.ResponseWriter, r *http.Request) {
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
	followersList, err := repo.Followers(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followersList)
}
// Stop follow and user
func ListFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	userID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	followersList, err := repo.Following(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followersList)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	
	tokenUserID, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if tokenUserID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("it's not yours user"))
		return
	}
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var passwords models.Password

	if err = json.Unmarshal(body, &passwords); err != nil {
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

	dbPassword, err := repo.GetUserPassword(userID)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = secure.Verify(dbPassword, passwords.Current); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}
	
	hash, err := secure.Hash(passwords.New)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	
	if err = repo.UpdatePassword(userID, string(hash)); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)

}