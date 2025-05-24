package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Load initializes the securecookie with the hash and block keys from the config.
func Load() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save register the information of the authenticated user in the cookie
func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}
	encrypted, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encrypted,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}
