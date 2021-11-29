package models

import "time"

type Todo struct {
	ID        int        `json:"id" gorm:"type:int;autoIncrement;not null"`
	Name      string     `json:"name" gorm:"type:varchar(80);not null"`
	IsDone    int        `json:"is_active" gorm:"type:int;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"type:timestamp;null"`
}
