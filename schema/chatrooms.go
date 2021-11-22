package schema

import "time"

//chatroom struct holds the chatroom details
type Chatroom struct {
	ID        uint       `gorm:"primaryKey", json:"id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	RoomName  string     `json:"room_name"`

	Users []*Users `gorm:"many2many:chatroom_users" json:"users"`
}

//join table
type ChatroomUsers struct {
	ChatroomID uint `"json:chatroom_id"`
	UsersID    uint `"json:users_id"`
}
