package repositories

import (
	"errors"
	"steveInterviewMod/config"
	"steveInterviewMod/models"
	"steveInterviewMod/utils"
	"time"
)

type LogRepository interface {
	GetByCondition(playerID *uint, action, startTime, endTime string, limit *int) ([]models.Log, error)
	Create(log *models.Log) error
}

type FileLogRepository struct{}

func UniLogRepository() LogRepository {
	if config.AppConfig.UseSQLDatabase {
		return nil
	}
	return &FileLogRepository{}
}

// File-based implementation
func (r *FileLogRepository) GetByCondition(playerID *uint, action, startTime, endTime string, limit *int) ([]models.Log, error) {
	var logs []models.Log
	err := utils.ReadFromFile(config.AppConfig.LogDataFilePath, &logs)
	if err != nil {
		return nil, err
	}

	filtered := []models.Log{}
	var start, end time.Time
	if startTime != "" {
		start, err = time.Parse("2006-01-02 15:04:05", startTime)
		if err != nil {
			return nil, errors.New("invalid start time format")
		}
	}
	if endTime != "" {
		end, err = time.Parse("2006-01-02 15:04:05", endTime)
		if err != nil {
			return nil, errors.New("invalid end time format")
		}
	}

	for _, log := range logs {
		logTime, _ := time.Parse("2006-01-02 15:04:05", log.Timestamp)
		if playerID != nil && log.PlayerID != *playerID {
			continue
		}
		if action != "" && string(log.Action) != action {
			continue
		}
		if startTime != "" && logTime.Before(start) {
			continue
		}
		if endTime != "" && logTime.After(end) {
			continue
		}
		filtered = append(filtered, log)
	}

	if limit != nil && len(filtered) > *limit {
		filtered = filtered[:*limit]
	}

	return filtered, nil
}

func (r *FileLogRepository) Create(log *models.Log) error {
	var logs []models.Log

	isPlayerExist, err := PlayerExistsFile(log.PlayerID, config.AppConfig.PlayerDataFilePath)
	if !isPlayerExist || err != nil {
		return errors.New("player does not exist")
	}

	err = utils.ReadFromFile(config.AppConfig.LogDataFilePath, &logs)
	if err != nil {
		return err
	}
	log.ID = uint(len(logs) + 1)
	log.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	logs = append(logs, *log)
	return utils.WriteToFile(config.AppConfig.LogDataFilePath, logs)
}
