package entity

import "time"

type Category struct {
	ID        string    `json:"id" gorm:"primaryKey;size:32"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug" gorm:"uniqueIndex:uniq_categories_slug;size:191"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
