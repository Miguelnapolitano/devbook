package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/models/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//creates a publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(body, &publication); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userID

	if err = publication.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPublicationsRepository(db)
	publication.ID, err = repo.Creates(publication)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

//List the publication on user feed
func ListPublications(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
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

	repo := repositories.NewPublicationsRepository(db)
	publications, err := repo.List(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)

}

//Retrieves a publication
func RetrievePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repo := repositories.NewPublicationsRepository(db)
	publication, err := repo.RetrieveByID(publicationID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	} 

	responses.JSON(w, http.StatusOK, publication)

}

//Updates a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repo := repositories.NewPublicationsRepository(db)
	dbPublication, err := repo.RetrieveByID(publicationID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	} 

	if dbPublication.AuthorID != tokenUserID {
		responses.Err(w, http.StatusForbidden, errors.New("it's not yor publication"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(body, &publication); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = publication.Prepare();  err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.Updates(publicationID, publication); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil) 

}

//Deletes a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.GetUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repo := repositories.NewPublicationsRepository(db)
	dbPublication, err := repo.RetrieveByID(publicationID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	} 

	if dbPublication.AuthorID != tokenUserID {
		responses.Err(w, http.StatusForbidden, errors.New("it's not yor publication"))
		return
	}

	if err = repo.Delete(publicationID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}


func ListUserPublications(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
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

	repo := repositories.NewPublicationsRepository(db)
	dbPublications, err := repo.ListByUser(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, dbPublications)
}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repo := repositories.NewPublicationsRepository(db)
	if err := repo.Like(publicationID); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	} 

	responses.JSON(w, http.StatusNoContent, nil)
}


func UnlikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repo := repositories.NewPublicationsRepository(db)
	if err := repo.Unlike(publicationID); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	} 

	responses.JSON(w, http.StatusNoContent, nil)
}


