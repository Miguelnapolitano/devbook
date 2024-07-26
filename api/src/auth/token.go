package auth

import (
	"api/src/config"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissions)
	fmt.Println(config.SecretKey)
	tokenString, err := token.SignedString([]byte(config.SecretKey))
        if err != nil {
                return "", err
        }

        return tokenString, nil
}
