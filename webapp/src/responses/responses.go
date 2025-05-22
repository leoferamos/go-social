package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorAPI represents a standard error response.
type ErrorAPI struct {
	Error string `json:"error"`
}

// JSON returns a JSON response to the client with the specified status code and data.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// HandleStatusCode is a helper function to handle different status codes and return appropriate responses.
func HandleStatusCode(w http.ResponseWriter, r *http.Response) {
	var err ErrorAPI
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
