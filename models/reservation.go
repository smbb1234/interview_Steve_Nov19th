package models

type Reservation struct {
	ID       uint   `json:"id"`
	RoomID   uint   `json:"room_id"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	PlayerID uint   `json:"player_id"`
}
