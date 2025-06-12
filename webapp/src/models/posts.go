package models

// Posts represents a post in the application.
type Posts struct {
	ID             uint64 `json:"id,omitempty"`
	Content        string `json:"content,omitempty"`
	AuthorID       uint64 `json:"author_id,omitempty"`
	AuthorUsername string `json:"author_username,omitempty"`
	Likes          uint64 `json:"likes"`
	CreatedAt      string `json:"created_at,omitempty"`
}
