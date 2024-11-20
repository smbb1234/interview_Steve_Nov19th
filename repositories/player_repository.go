package repositories

import (
	"errors"
	"steveInterviewMod/config"
	"steveInterviewMod/models"
	"steveInterviewMod/utils"
	"strconv"
)

type PlayerRepository interface {
	GetAll() ([]models.Player, error)
	GetByID(id string) (*models.Player, error)
	Create(player *models.Player) error
	Update(player models.Player) error
	Delete(id string) error
}

type FilePlayerRepository struct{}

func UniPlayerRepository() PlayerRepository {
	if config.AppConfig.UseSQLDatabase {
		return nil
	}
	return &FilePlayerRepository{}
}

// File-based implementation
func (r *FilePlayerRepository) GetAll() ([]models.Player, error) {
	var players []models.Player
	err := utils.ReadFromFile(config.AppConfig.PlayerDataFilePath, &players)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (r *FilePlayerRepository) GetByID(id string) (*models.Player, error) {
	players, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	playerID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid player ID")
	}

	for _, player := range players {
		if player.ID == uint(playerID) {
			return &player, nil
		}
	}
	return nil, errors.New("player not found")
}

func (r *FilePlayerRepository) Create(player *models.Player) error {
	players, err := r.GetAll()
	if err != nil {
		return err
	}

	player.ID = uint(len(players) + 1)
	players = append(players, *player)
	return utils.WriteToFile(config.AppConfig.PlayerDataFilePath, players)
}

func (r *FilePlayerRepository) Update(updatedPlayer models.Player) error {
	players, err := r.GetAll()
	if err != nil {
		return err
	}

	for i, player := range players {
		if player.ID == updatedPlayer.ID {
			players[i] = updatedPlayer
			return utils.WriteToFile(config.AppConfig.PlayerDataFilePath, players)
		}
	}
	return errors.New("player not found")
}

func (r *FilePlayerRepository) Delete(id string) error {
	players, err := r.GetAll()
	if err != nil {
		return err
	}

	playerID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid player id")
	}

	for i, player := range players {
		if player.ID == uint(playerID) {
			players = append(players[:i], players[i+1:]...)
			return utils.WriteToFile(config.AppConfig.PlayerDataFilePath, players)
		}
	}
	return errors.New("player not found")
}
