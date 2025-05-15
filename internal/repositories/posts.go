package repositories

import (
	"database/sql"
	"go_social/internal/models"
)

// Posts represents the repository for posts.
type Posts struct {
	db *sql.DB
}

// NewPostsRepository creates a new Posts repository.
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// CreatePost creates a new post in the database.
func (r Posts) CreatePost(post models.Posts) (uint64, error) {
	statement, err := r.db.Prepare(
		"INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertId), nil
}
