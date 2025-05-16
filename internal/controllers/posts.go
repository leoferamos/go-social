package controllers

import (
	"encoding/json"
	"go_social/internal/auth"
	"go_social/internal/db"
	"go_social/internal/models"
	"go_social/internal/repositories"
	"go_social/internal/responses"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost handles the creation of a new post.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.JSONError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Posts
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err = post.Prepare(); err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post.ID, err = repository.CreatePost(post)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, post)
}

// GetPost handles the retrieval of a post by its ID.
func GetPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewPostsRepository(db)
	post, err := repository.GetPostByID(postID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, post)
}

// GetPosts Gets the posts of people the user follows and their own posts.
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
	db, err := db.Connect()
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewPostsRepository(db)
	posts, err := repository.GetPosts(userID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

// UpdatePost Updates a post.
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a post
}

// DeletePost Deletes a post.
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a post
}
