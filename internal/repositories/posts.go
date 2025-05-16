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

// GetPost retrieves a post by its ID from the database.
func (r Posts) GetPostByID(id uint64) (models.Posts, error) {
	rows, err := r.db.Query(
		`SELECT p.id, p.title, p.content, p.author_id, u.username, p.likes, p.created_at
      FROM posts p
      INNER JOIN users u ON u.id = p.author_id
      WHERE p.id = ?`,
		id,
	)
	if err != nil {
		return models.Posts{}, err
	}
	defer rows.Close()
	var post models.Posts
	if rows.Next() {
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorUsername,
			&post.Likes,
			&post.CreatedAt,
		); err != nil {
			return models.Posts{}, err
		}
	}
	return post, nil
}

// GetPosts Gets the posts of people the user follows and their own posts.
func (r Posts) GetPosts(userID uint64) ([]models.Posts, error) {
	rows, err := r.db.Query(
		`SELECT p.id, p.title, p.content, p.author_id, u.username, p.likes, p.created_at
    FROM posts p
    INNER JOIN users u ON u.id = p.author_id
    LEFT JOIN followers f ON f.user_id = p.author_id AND f.follower_id = ?
    WHERE f.follower_id IS NOT NULL OR p.author_id = ?
    ORDER BY 1 DESC`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Posts
	for rows.Next() {
		var post models.Posts
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorUsername,
			&post.Likes,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// UpdatePost updates a post in the database.
func (r Posts) UpdatePost(postID uint64, post models.Posts) error {
	statement, err := r.db.Prepare(
		"UPDATE posts SET title = ?, content = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}
	return nil
}

// DeletePost deletes a post from the database.
func (r Posts) DeletePost(postID uint64) error {
	statement, err := r.db.Prepare(
		"DELETE FROM posts WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(postID); err != nil {
		return err
	}
	return nil
}
