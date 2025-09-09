package entity

import "time"

type ContentStatus string

const (
	StatusDraft     ContentStatus = "draft"
	StatusPublished ContentStatus = "published"
)

type Content struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Slug        string        `json:"slug"`
	Body        string        `json:"body"`
	Excerpt     string        `json:"excerpt"`
	Image       string        `json:"image"`
	Status      ContentStatus `json:"status"`
	AuthorID    string        `json:"author_id"`
	CategoryID  string        `json:"category_id"`
	Tags        []string      `json:"tags"`
	PublishedAt *time.Time    `json:"published_at,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
