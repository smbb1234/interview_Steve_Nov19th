package services

import (
	"errors"
	"steveInterviewMod/models"
	"steveInterviewMod/repositories"
)

var levelRepo repositories.LevelRepository = repositories.UniLevelRepository()

func GetAllLevels() ([]models.Level, error) {
	return levelRepo.GetAllLevels()
}

func CreateLevel(request map[string]interface{}) (*models.Level, error) {

	name, ok := request["name"].(string)
	if !ok || name == "" {
		return nil, errors.New("level name is required")
	}
	level := models.Level{
		Name: name,
	}
	err := levelRepo.CreateLevel(&level)
	if err != nil {
		return nil, err
	}
	return &level, nil
}
