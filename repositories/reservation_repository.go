package repositories

import (
	"errors"
	"steveInterviewMod/config"
	"steveInterviewMod/models"
	"steveInterviewMod/utils"
)

type ReservationRepository interface {
	GetByCondition(roomID *uint, date string, limit *int) ([]models.Reservation, error)
	Create(reservation *models.Reservation) error
}

type FileReservationRepository struct{}

func UniReservationRepository() ReservationRepository {
	if config.AppConfig.UseSQLDatabase {
		return nil
	}
	return &FileReservationRepository{}
}

// File-based implementation
func (r *FileReservationRepository) GetByCondition(roomID *uint, date string, limit *int) ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := utils.ReadFromFile(config.AppConfig.ReservationDataFilePath, &reservations)
	if err != nil {
		return nil, err
	}

	filtered := []models.Reservation{}
	for _, reservation := range reservations {
		if roomID != nil && reservation.RoomID != *roomID {
			continue
		}
		if date != "" && reservation.Date != date {
			continue
		}
		filtered = append(filtered, reservation)
	}

	if limit != nil && len(filtered) > *limit {
		filtered = filtered[:*limit]
	}

	return filtered, nil
}

func (r *FileReservationRepository) Create(reservation *models.Reservation) error {
	var reservations []models.Reservation

	isRoomExist, err := RoomExistsFile(reservation.RoomID, config.AppConfig.RoomDataFilePath)
	if !isRoomExist || err != nil {
		return errors.New("room does not exist")
	}
	isPlayerExist, err := PlayerExistsFile(reservation.PlayerID, config.AppConfig.PlayerDataFilePath)
	if !isPlayerExist || err != nil {
		return errors.New("player does not exist")
	}

	err = utils.ReadFromFile(config.AppConfig.ReservationDataFilePath, &reservations)
	if err != nil {
		return err
	}

	reservation.ID = uint(len(reservations) + 1)
	reservations = append(reservations, *reservation)
	return utils.WriteToFile(config.AppConfig.ReservationDataFilePath, reservations)
}
