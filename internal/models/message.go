package models

import "time"

type Message struct {
	ID        string    `gorm:"type:uuid;primary_key" json:"message_id"`
	Content   string    `gorm:"column:content; type:text; not null" json:"content"`
	RoomID    string    `gorm:"type:uuid;not null" json:"room_id"`
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
}
