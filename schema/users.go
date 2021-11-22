package schema

import "time"

//users struct holds the user information
type Users struct {
	ID        uint       `gorm:"primaryKey", json:"id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	// ChatRooms []*Chatroom `gorm:"many2many:chatroom_users"`
}
