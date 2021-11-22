package types

type ChatroomReq struct {
	RoomName string `json:"room_name"`
	UserIDs  []uint `json:"user_ids"`
}
