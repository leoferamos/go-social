package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID              uint64    `json:"id,omitempty"`
	Content         string    `json:"content,omitempty"`
	AuthorID        uint64    `json:"author_id,omitempty"`
	AuthorUsername  string    `json:"author_username,omitempty"`
	AuthorName      string    `json:"author_name,omitempty"`
	AuthorAvatarURL string    `json:"author_avatar_url,omitempty"`
	Likes           uint64    `json:"likes"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	LikedByMe       bool      `json:"liked_by_me"`
}

// Prepare calls the methods to format and validate the post data.
func (p *Post) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}
	p.format()
	return nil
}
func (p *Post) validate() error {
	if p.Content == "" {
		return errors.New("content is required")
	}
	return nil
}
func (p *Post) format() {
	p.Content = strings.TrimSpace(p.Content)
}
