package entity

import (
	"time"

	"gorm.io/datatypes"
)

type ContentStatus string

const (
	StatusDraft     ContentStatus = "draft"
	StatusPublished ContentStatus = "published"
)

type Content struct {
	ID          string         `json:"id" gorm:"primaryKey;size:32"`
	Title       string         `json:"title"`
	Slug        string         `json:"slug" gorm:"uniqueIndex:uniq_contents_slug;size:191"`
	Body        string         `json:"body" gorm:"type:text"`
	Excerpt     string         `json:"excerpt" gorm:"type:text"`
	Image       string         `json:"image" gorm:"type:text"`
	Status      ContentStatus  `json:"status" gorm:"index:idx_contents_status;size:32"`
	AuthorID    string         `json:"author_id" gorm:"index:idx_contents_author;size:32"`
	CategoryID  string         `json:"category_id" gorm:"index:idx_contents_category;size:32"`
	Tags        datatypes.JSON `json:"tags" gorm:"type:jsonb"`
	PublishedAt *time.Time     `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
