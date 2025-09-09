package entity

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;size:32"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"uniqueIndex:uniq_users_email;size:191"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
