package services

import (
	"errors"
	"steveInterviewMod/models"
	"steveInterviewMod/repositories"
)

var logRepo repositories.LogRepository = repositories.UniLogRepository()

func GetLogsByCondition(playerID *uint, action, startTime, endTime string, limit *int) ([]models.Log, error) {

	return logRepo.GetByCondition(playerID, action, startTime, endTime, limit)

}

func CreateLog(request map[string]interface{}) (*models.Log, error) {

	playerId, ok := request["player_id"].(float64)
	if !ok {
		return nil, errors.New("player information is required")
	}
	actionStr, ok := request["action"].(string)
	if !ok {
		return nil, errors.New("action id is required")
	}
	action, err := models.ToAction(actionStr)
	if err != nil {
		return nil, err
	}
	details, _ := request["details"].(string)

	log := models.Log{
		PlayerID: uint(playerId),
		Action:   action,
		Details:  details,
	}

	err = logRepo.Create(&log)
	if err != nil {
		return nil, err
	}
	return &log, nil
}
