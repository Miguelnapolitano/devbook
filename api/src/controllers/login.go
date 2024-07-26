package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/models/repositories"
	"api/src/responses"
	"api/src/secure"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
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

	dbUser, err := repo.FindByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := secure.Verify(dbUser.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(dbUser.ID)
	if err != nil {
			fmt.Println("Erro ao gerar token:", err)
			return
	}

	
	fmt.Println("Token gerado:", token)

}