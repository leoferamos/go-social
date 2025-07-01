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
func (r Posts) CreatePost(post models.Post) (uint64, error) {
	statement, err := r.db.Prepare(
		"INSERT INTO posts (content, author_id) VALUES (?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(post.Content, post.AuthorID)
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
func (r Posts) GetPostByID(id uint64, userID uint64) (models.Post, error) {
	rows, err := r.db.Query(
		`SELECT p.id, p.content, p.author_id, u.username, u.name,
			COUNT(pl.user_id) AS likes, p.created_at,
			CASE WHEN EXISTS (
				SELECT 1 FROM post_likes pl2 WHERE pl2.post_id = p.id AND pl2.user_id = ?
			) THEN 1 ELSE 0 END AS liked_by_me
		FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		LEFT JOIN post_likes pl ON pl.post_id = p.id
		WHERE p.id = ?
		GROUP BY p.id, p.content, p.author_id, u.username, u.name, p.created_at
		`,
		userID, id,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()
	var post models.Post
	if rows.Next() {
		if err = rows.Scan(
			&post.ID,
			&post.Content,
			&post.AuthorID,
			&post.AuthorUsername,
			&post.AuthorName,
			&post.Likes,
			&post.CreatedAt,
			&post.LikedByMe,
		); err != nil {
			return models.Post{}, err
		}
	}
	return post, nil
}

// GetPosts Gets the posts of people the user follows and their own posts.
func (r Posts) GetPosts(userID uint64) ([]models.Post, error) {
	rows, err := r.db.Query(
		`SELECT p.id, p.content, p.author_id, u.username, u.name,
			COUNT(pl.user_id) AS likes, p.created_at,
			CASE WHEN EXISTS (
				SELECT 1 FROM post_likes pl2 WHERE pl2.post_id = p.id AND pl2.user_id = ?
			) THEN 1 ELSE 0 END AS liked_by_me
		FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		LEFT JOIN followers f ON f.user_id = p.author_id AND f.follower_id = ?
		LEFT JOIN post_likes pl ON pl.post_id = p.id
		WHERE f.follower_id IS NOT NULL OR p.author_id = ?
		GROUP BY p.id, p.content, p.author_id, u.username, u.name, p.created_at
		ORDER BY p.created_at DESC`,
		userID, userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Content,
			&post.AuthorID,
			&post.AuthorUsername,
			&post.AuthorName,
			&post.Likes,
			&post.CreatedAt,
			&post.LikedByMe,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// UpdatePost updates a post in the database.
func (r Posts) UpdatePost(postID uint64, post models.Post) error {
	statement, err := r.db.Prepare(
		"UPDATE posts SET content = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(post.Content, postID); err != nil {
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

// GetUserPosts Gets the posts of a user.
func (r Posts) GetUserPosts(targetUserID uint64, userID uint64) ([]models.Post, error) {
	rows, err := r.db.Query(
		`SELECT p.id, p.content, p.author_id, u.username, u.name,
			COUNT(pl.user_id) AS likes, p.created_at,
			CASE WHEN EXISTS (
				SELECT 1 FROM post_likes pl2 WHERE pl2.post_id = p.id AND pl2.user_id = ?
			) THEN 1 ELSE 0 END AS liked_by_me
		FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		LEFT JOIN post_likes pl ON pl.post_id = p.id
		WHERE p.author_id = ?
		GROUP BY p.id, p.content, p.author_id, u.username, u.name, p.created_at
		ORDER BY p.created_at DESC`,
		userID, targetUserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Content,
			&post.AuthorID,
			&post.AuthorUsername,
			&post.AuthorName,
			&post.Likes,
			&post.CreatedAt,
			&post.LikedByMe,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// LikePost increments the like count of a post in the database.
func (r Posts) LikePost(userID, postID uint64) error {
	_, err := r.db.Exec(
		"INSERT IGNORE INTO post_likes (user_id, post_id) VALUES (?, ?)",
		userID, postID,
	)
	return err
}

// UnlikePost decrements the like count of a post in the database.
func (r Posts) UnlikePost(userID, postID uint64) error {
	_, err := r.db.Exec(
		"DELETE FROM post_likes WHERE user_id = ? AND post_id = ?",
		userID, postID,
	)
	return err
}
