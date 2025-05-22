package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a JSON response to the client with the specified status code and data.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
