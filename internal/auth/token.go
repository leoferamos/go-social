package auth

import (
	"go_social/config"
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
