package repositories

import (
	"api/src/models"
	"database/sql"
)

//represents a user repo
type users struct {
	db *sql.DB
}

//Creates a user repo
func NewUserRepository (db *sql.DB) *users {
	return &users{db}
} 

//Inserts an user at DB
func (repo users) Creates(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()
	
	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}