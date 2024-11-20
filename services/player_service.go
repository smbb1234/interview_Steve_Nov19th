package services

import (
	"errors"
	"steveInterviewMod/models"
	"steveInterviewMod/repositories"
)

var playerRepo repositories.PlayerRepository = repositories.UniPlayerRepository()

func GetAllPlayers() ([]models.Player, error) {
	return playerRepo.GetAll()
}

func CreatePlayer(request map[string]interface{}) (*models.Player, error) {

	name, ok := request["name"].(string)
	if !ok || name == "" {
		return nil, errors.New("player name is required")
	}
	levelID, ok := request["level_id"].(float64)
	if !ok {
		return nil, errors.New("level is required")
	}
	others, _ := request["others"].(string)

	player := models.Player{
		Name:    name,
		LevelID: uint(levelID),
		Others:  others,
	}
	err := playerRepo.Create(&player)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func GetPlayerByID(id string) (*models.Player, error) {
	return playerRepo.GetByID(id)
}

func UpdatePlayer(id string, request map[string]interface{}) (*models.Player, error) {
	player, err := playerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name, ok := request["name"].(string); ok {
		player.Name = name
	}

	if levelID, ok := request["level_id"].(float64); ok {
		player.LevelID = uint(levelID)
	}
	if others, ok := request["others"].(string); ok {
		player.Others = others
	}

	err = playerRepo.Update(*player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func DeletePlayer(id string) error {
	return playerRepo.Delete(id)
}
