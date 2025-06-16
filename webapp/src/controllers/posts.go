package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"webapp/src/requests"
	"webapp/src/responses"
)

// CreatePost calls the API to create a new post.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Invalid JSON"})
		return
	}
	if body.Content == "" {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Content is required"})
		return
	}

	post, err := json.Marshal(map[string]string{
		"content": body.Content,
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}
	apiURL = fmt.Sprintf("%s/posts", apiURL)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, apiURL, bytes.NewBuffer(post))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	log.Printf("[DEBUG] Status de resposta da API: %d", response.StatusCode)
	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	// Repasse o corpo da resposta da API (post criado) para o frontend
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if _, err := io.Copy(w, response.Body); err != nil {
		log.Printf("[DEBUG] Erro ao copiar resposta da API: %v", err)
	}
}
