package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON is a helper function to write JSON responses to the client.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println("error encoding JSON response:", err)
		}
	}
}

// JSONError is a helper function to write JSON error responses to the client.
func JSONError(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
