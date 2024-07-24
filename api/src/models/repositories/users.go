package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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
		"insert into users (name, nick, email, password), values(?, ?, ?, ?)",
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

//all users with name or nick received
func (repo users) List(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repo.db.Query(
		"select id, name, nick, email, created_at from users where name like ? or nick like ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err 
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Created_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}