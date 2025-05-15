package models

import (
	"errors"
	"strings"
	"time"
)

type Posts struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"author_id,omitempty"`
	AuthorUsername string    `json:"author_username,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

// Prepare calls the methods to format and validate the post data.
func (p *Posts) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}
	p.format()
	return nil
}
func (p *Posts) validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}
	if p.Content == "" {
		return errors.New("content is required")
	}
	return nil
}
func (p *Posts) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}
