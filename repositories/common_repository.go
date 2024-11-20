package repositories

import (
	"steveInterviewMod/models"
	"steveInterviewMod/utils"
)

// File-based implementation
func RoomExistsFile(roomID uint, filePath string) (bool, error) {
	var rooms []models.Room
	err := utils.ReadFromFile(filePath, &rooms)
	if err != nil {
		return false, err
	}
	for _, room := range rooms {
		if room.ID == roomID {
			return true, nil
		}
	}
	return false, nil
}

func PlayerExistsFile(playerID uint, filePath string) (bool, error) {
	var players []models.Player
	err := utils.ReadFromFile(filePath, &players)
	if err != nil {
		return false, err
	}
	for _, player := range players {
		if player.ID == playerID {
			return true, nil
		}
	}
	return false, nil
}
