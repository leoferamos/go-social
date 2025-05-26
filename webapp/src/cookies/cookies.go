package cookies

import (
	"errors"
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Load initializes the securecookie instance with the hash and block keys from the config.
func Load() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save encrypts the user ID and token, then sets them as a cookie in the HTTP response.
func Save(w http.ResponseWriter, ID, token string) error {
	if s == nil {
		return errors.New("securecookie not initialized, call Load() first")
	}

	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encrypted, err := s.Encode("auth_data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_data",
		Value:    encrypted,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
		MaxAge:   86400,
	})

	return nil
}

// Read return the value of the cookie as a map[string]string.
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("auth_data")
	if err != nil {
		return nil, err
	}

	value := make(map[string]string)

	if err = s.Decode("auth_data", cookie.Value, &value); err != nil {
		return nil, err
	}

	return value, nil
}
