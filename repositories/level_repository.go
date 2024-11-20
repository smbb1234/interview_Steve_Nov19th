package repositories

import (
	"steveInterviewMod/config"
	"steveInterviewMod/models"
	"steveInterviewMod/utils"
)

type LevelRepository interface {
	GetAllLevels() ([]models.Level, error)
	CreateLevel(level *models.Level) error
}

type FileLevelRepository struct{}

func UniLevelRepository() LevelRepository {
	if config.AppConfig.UseSQLDatabase {
		return nil
	}
	return &FileLevelRepository{}
}

// File-based implementation
func (r *FileLevelRepository) GetAllLevels() ([]models.Level, error) {
	var levels []models.Level
	err := utils.ReadFromFile(config.AppConfig.LevelDataFilePath, &levels)
	if err != nil {
		return nil, err
	}
	return levels, nil
}

func (r *FileLevelRepository) CreateLevel(level *models.Level) error {
	levels, err := r.GetAllLevels()
	if err != nil {
		return err
	}

	level.ID = uint(len(levels) + 1)
	levels = append(levels, *level)
	return utils.WriteToFile(config.AppConfig.LevelDataFilePath, levels)
}
