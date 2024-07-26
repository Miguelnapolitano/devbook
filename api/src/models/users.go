package models

import (
	"api/src/secure"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user using a social media
type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}

// Call validate and format user
func (user *User) Prepare(method string) error {
	if err := user.validate(method); err != nil {
		return err
	}

	if err := user.format(method); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(method string) error {
	if user.Name == "" {
		return errors.New("the name is required and blank not allowed")
	}
	if user.Nick == "" {
		return errors.New("the nick is required and blank not allowed")
	}
	if user.Email == "" {
		return errors.New("the email is required and blank not allowed")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return err
	}

	if method == "POST" && user.Password == "" {
		return errors.New("the password is required and blank not allowed")
	}
	return nil
}

func (user *User) format(method string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if method == "POST" {
		hash, err := secure.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hash)

	}

	return nil
}
