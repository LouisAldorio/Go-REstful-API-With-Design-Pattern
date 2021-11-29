package models

import "time"

type User struct {
	ID        int        `json:"id" gorm:"type:int;autoIncrement;not null"`
	Name      string     `json:"name" gorm:"type:varchar(80);not null"`
	Email     string     `json:"email" gorm:"type:varchar(80);not null"`
	Password  string     `json:"-" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"type:timestamp;null"`
}
