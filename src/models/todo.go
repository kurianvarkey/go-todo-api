package models

import (
	"time"

	"gorm.io/datatypes"
)

type Todo struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	UserId        int            `json:"user_id"`
	Title         string         `json:"title" gorm:"size:100"`
	Description   string         `json:"description" gorm:"size:500"`
	TaskDate      datatypes.Date `json:"task_date"`
	IsCompleted   bool           `json:"is_completed"`
	CompletedDate time.Time      `json:"completed_date"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     time.Time      `json:"deleted_at"`
	Deleted       bool           `json:"deleted"`
}

/* type Product struct {
	gorm.Model
	Code  string
	Price uint
} */
