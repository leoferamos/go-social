package auth

import (
	"errors"
	"fmt"
	"go_social/config"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken generates a JWT token for the given user ID.
func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken checks if the provided token is valid.
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnKeyFunc)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return nil
	}
	return errors.New("Invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.SecretKey), nil
}
