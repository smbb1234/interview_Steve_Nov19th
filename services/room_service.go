package services

import (
	"errors"
	"steveInterviewMod/models"
	"steveInterviewMod/repositories"
)

var roomRepo repositories.RoomRepository = repositories.UniRoomRepository()

func GetAllRooms() ([]models.Room, error) {
	return roomRepo.GetAll()
}

func CreateRoom(request map[string]interface{}) (*models.Room, error) {

	name, ok := request["name"].(string)
	if !ok || name == "" {
		return nil, errors.New("room name is required")
	}
	status, ok := request["status"].(string)
	if !ok || status == "" {
		return nil, errors.New("room status is required")
	}
	description, _ := request["description"].(string)

	room := models.Room{
		Name:        name,
		Description: description,
		Status:      status,
	}
	err := roomRepo.Create(&room)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func GetRoomByID(id string) (*models.Room, error) {
	return roomRepo.GetByID(id)
}

func UpdateRoom(id string, request map[string]interface{}) (*models.Room, error) {
	room, err := roomRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name, ok := request["name"].(string); ok {
		room.Name = name
	}
	if status, ok := request["status"].(string); ok {
		room.Status = status
	}
	if description, ok := request["description"].(string); ok {
		room.Description = description
	}

	err = roomRepo.Update(*room)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func DeleteRoom(id string) error {
	return roomRepo.Delete(id)
}
