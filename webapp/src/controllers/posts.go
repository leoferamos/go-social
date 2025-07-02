package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
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

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	var createdPost models.Post
	if err := json.NewDecoder(response.Body).Decode(&createdPost); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to parse created post"})
		return
	}

	if createdPost.AuthorAvatarURL == "" {
		createdPost.AuthorAvatarURL = "/assets/img/avatar-placeholder.png"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var buf bytes.Buffer
	err = utils.ExecuteTemplate(&buf, "post-with-permission", createdPost)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Failed to render post template"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

// LikePost calls the API to like a post.
func LikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]
	if postID == "" {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Post ID is required"})
		return
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}
	apiURL = fmt.Sprintf("%s/posts/%s/like", apiURL, postID)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, apiURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if _, err := io.Copy(w, response.Body); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
}

// UnlikePost calls the API to unlike a post.
func UnlikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]
	if postID == "" {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Post ID is required"})
		return
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}
	apiURL = fmt.Sprintf("%s/posts/%s/unlike", apiURL, postID)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, apiURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	if _, err := io.Copy(w, response.Body); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
}

// DeletePost calls the API to delete a post.
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]
	if postID == "" {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: "Post ID is required"})
		return
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://api:5000"
	}
	apiURL = fmt.Sprintf("%s/posts/%s", apiURL, postID)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodDelete, apiURL, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCode(w, response)
		return
	}

	w.WriteHeader(response.StatusCode)
}
