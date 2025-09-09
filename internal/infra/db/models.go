package db

import (
	"time"
)

// GORM models (kept separate from entities to avoid leaking gorm tags upward)
type User struct {
	ID        string `gorm:"primaryKey;size:32"`
	Name      string
	Email     string `gorm:"uniqueIndex;size:191"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Category struct {
	ID        string `gorm:"primaryKey;size:32"`
	Name      string
	Slug      string `gorm:"uniqueIndex;size:191"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Content struct {
	ID          string `gorm:"primaryKey;size:32"`
	Title       string
	Slug        string `gorm:"uniqueIndex;size:191"`
	Body        string `gorm:"type:longtext"`
	Status      string `gorm:"index;size:32"`
	AuthorID    string `gorm:"index;size:32"`
	CategoryID  string `gorm:"index;size:32"`
	PublishedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
