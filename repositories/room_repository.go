package repositories

import (
	"errors"
	"steveInterviewMod/config"
	"steveInterviewMod/models"
	"steveInterviewMod/utils"
	"strconv"
)

type RoomRepository interface {
	GetAll() ([]models.Room, error)
	GetByID(id string) (*models.Room, error)
	Create(room *models.Room) error
	Update(room models.Room) error
	Delete(id string) error
}

type FileRoomRepository struct{}

func UniRoomRepository() RoomRepository {
	if config.AppConfig.UseSQLDatabase {
		return nil
	}
	return &FileRoomRepository{}
}

// File-based implementation
func (r *FileRoomRepository) GetAll() ([]models.Room, error) {
	var rooms []models.Room
	err := utils.ReadFromFile(config.AppConfig.RoomDataFilePath, &rooms)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *FileRoomRepository) GetByID(id string) (*models.Room, error) {
	rooms, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	roomID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid room id")
	}

	for _, room := range rooms {
		if room.ID == uint(roomID) {
			return &room, nil
		}
	}
	return nil, errors.New("room not found")
}

func (r *FileRoomRepository) Create(room *models.Room) error {
	rooms, err := r.GetAll()
	if err != nil {
		return err
	}
	room.ID = uint(len(rooms) + 1)
	rooms = append(rooms, *room)
	return utils.WriteToFile(config.AppConfig.RoomDataFilePath, rooms)
}

func (r *FileRoomRepository) Update(updatedRoom models.Room) error {
	rooms, err := r.GetAll()
	if err != nil {
		return err
	}
	for i, room := range rooms {
		if room.ID == updatedRoom.ID {
			rooms[i] = updatedRoom
			return utils.WriteToFile(config.AppConfig.RoomDataFilePath, rooms)
		}
	}
	return errors.New("room not found")
}

func (r *FileRoomRepository) Delete(id string) error {
	rooms, err := r.GetAll()
	if err != nil {
		return err
	}

	roomID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid room id")
	}

	for i, room := range rooms {
		if room.ID == uint(roomID) {
			rooms = append(rooms[:i], rooms[i+1:]...)
			return utils.WriteToFile(config.AppConfig.RoomDataFilePath, rooms)
		}
	}
	return errors.New("room not found")
}
