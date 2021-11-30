package models

import "time"

type Todo struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	IsDone    bool       `json:"is_done" gorm:"column:is_done"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	UserID    int       `json:"user_id" gorm:"column:user_id"`

	// User User `json:"user" gorm:"-"`
}
