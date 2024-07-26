package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// represents a user repo
type users struct {
	db *sql.DB
}

// Creates a user repo
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Inserts an user at DB
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

// all users with name or nick received
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
		err := rows.Scan(
			&user.ID, 
			&user.Name, 
			&user.Nick, 
			&user.Email, 
			&user.Created_at,
		)
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

// User with received ID
func (repo users) RetrieveUser(userID uint64) (models.User, error) {
	rows, err := repo.db.Query(
		"select id, name, nick, email, created_at from users where id = ?",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID, 
			&user.Name, 
			&user.Nick, 
			&user.Email, 
			&user.Created_at,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

// Updates User with received ID
func (repo users) UpdateUser(userID uint64, user models.User) error {
	statement, err := repo.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}

// Deletes User with received ID
func (repo users) DeleteUser(userID uint64) error {
	statement, err := repo.db.Prepare(
		"delete from users where id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

// Find an User by received email
func (repo users) FindByEmail(email string) (models.User, error){
	row, err := repo.db.Query(
		"select id, password from users where email = ?", 
		email,
	)
	if err != nil {
		return models.User{}, err
	}

	defer row.Close()

	var user models.User

	if row.Next(){
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}