package schema

import "time"

//message structs have chatroom message
type Messages struct {
	ID        uint       `gorm:"primaryKey", json:"id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Content   string     `json:"content"`
	RoomID    uint       `json:"room_id"`
	UserID    uint       `json:"user_id"`
}
