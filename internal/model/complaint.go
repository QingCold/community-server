package model

import "time"

type Complaint struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	User      string    `json:"user"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
