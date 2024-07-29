package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(r *http.Request) error {

	tokenStr := extractToken(r)

	token, err := jwt.Parse(tokenStr, returnsVerifyKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalid")
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}

func returnsVerifyKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("sign method unexpected: %v",
			token.Header["alg"],
		)
	}

	return config.SecretKey, nil
}

func GetUserID(r *http.Request) (uint64, error) {
	tokenStr := extractToken(r)

	token, err := jwt.Parse(tokenStr, returnsVerifyKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("invalid token")
}
