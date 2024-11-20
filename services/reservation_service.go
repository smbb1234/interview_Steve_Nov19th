package services

import (
	"errors"
	"steveInterviewMod/models"
	"steveInterviewMod/repositories"
)

var reservationRepo repositories.ReservationRepository = repositories.UniReservationRepository()

func GetReservations(roomID *uint, date string, limit *int) ([]models.Reservation, error) {
	return reservationRepo.GetByCondition(roomID, date, limit)
}

func CreateReservation(request map[string]interface{}) (*models.Reservation, error) {

	roomId, ok := request["room_id"].(float64)
	if !ok {
		return nil, errors.New("room id is required")
	}
	date, ok := request["date"].(string)
	if !ok || date == "" {
		return nil, errors.New("date is required")
	}
	time, ok := request["time"].(string)
	if !ok || time == "" {
		return nil, errors.New("time is required")
	}
	playerId, ok := request["player_id"].(float64)
	if !ok {
		return nil, errors.New("player information is required")
	}

	reservation := models.Reservation{
		RoomID:   uint(roomId),
		Date:     date,
		Time:     time,
		PlayerID: uint(playerId),
	}

	err := reservationRepo.Create(&reservation)
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}
