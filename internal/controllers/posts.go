package controllers

import (
	"encoding/json"
	"errors"
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

	createdPost, err := repository.GetPostByID(post.ID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, createdPost)
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
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
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
	postOnDB, err := repository.GetPostByID(postID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	if postOnDB.AuthorID != userID {
		responses.JSONError(w, http.StatusForbidden, errors.New("you are not the author of this post"))
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
	if err = post.Prepare(); err != nil {
		responses.JSONError(w, http.StatusBadRequest, err)
		return
	}
	if err = repository.UpdatePost(postID, post); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePost Deletes a post.
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
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
	postOnDB, err := repository.GetPostByID(postID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	if postOnDB.AuthorID != userID {
		responses.JSONError(w, http.StatusForbidden, errors.New("you are not the author of this post"))
		return
	}
	if err = repository.DeletePost(postID); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// GetUserPosts Gets the posts of a user.
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
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
	posts, err := repository.GetUserPosts(userID)
	if err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

// LikePost Likes a post.
func LikePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
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
	if err = repository.LikePost(userID, postID); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// UnlikePost Unlikes a post.
func UnlikePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.JSONError(w, http.StatusUnauthorized, err)
		return
	}
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
	if err = repository.UnlikePost(userID, postID); err != nil {
		responses.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
