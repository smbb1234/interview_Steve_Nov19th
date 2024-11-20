package repositories

import (
	"errors"
	"steveInterviewMod/config"
	"steveInterviewMod/models"
	"steveInterviewMod/utils"
	"time"
)

type ChallengeRepository interface {
	GetRecentChallenges() ([]models.Challenge, error)
	Create(challenge *models.Challenge) error
	GetLastByPlayer(playerID uint) (*models.Challenge, error)
	GetPendingChallenges() ([]models.Challenge, error)
	UpdateStatus(challenge models.Challenge, status models.ChallengeStatus) error
}

type FileChallengeRepository struct{}

func UniChallengeRepository() ChallengeRepository {
	if config.AppConfig.UseSQLDatabase {
		return nil
	}
	return &FileChallengeRepository{}
}

// File-based implementation
func (r *FileChallengeRepository) GetRecentChallenges() ([]models.Challenge, error) {
	var challenges []models.Challenge
	err := utils.ReadFromFile(config.AppConfig.ChallengesDataFilePath, &challenges)
	if err != nil {
		return nil, err
	}
	return challenges, nil
}

func (r *FileChallengeRepository) GetLastByPlayer(playerID uint) (*models.Challenge, error) {
	var challenges []models.Challenge
	err := utils.ReadFromFile(config.AppConfig.ChallengesDataFilePath, &challenges)
	if err != nil {
		return nil, err
	}

	var lastChallenge *models.Challenge
	var lastTime time.Time

	for _, challenge := range challenges {
		if challenge.PlayerID == playerID {
			challengeTime, err := time.Parse("2006-01-02 15:04:05", challenge.Timestamp)
			if err != nil {
				continue
			}
			if lastChallenge == nil || challengeTime.After(lastTime) {
				lastChallenge = &challenge
				lastTime = challengeTime
			}
		}
	}

	return lastChallenge, nil
}

func (r *FileChallengeRepository) Create(challenge *models.Challenge) error {
	var challenges []models.Challenge

	isPlayerExist, err := PlayerExistsFile(challenge.PlayerID, config.AppConfig.PlayerDataFilePath)
	if !isPlayerExist || err != nil {
		return errors.New("player does not exist")
	}

	lastChallenge, err := r.GetLastByPlayer(challenge.PlayerID)
	if err != nil {
		return err
	}
	if lastChallenge != nil {
		lastTime, err := time.Parse("2006-01-02 15:04:05", lastChallenge.Timestamp)
		if err != nil {
			return err
		}
		if time.Since(lastTime) < time.Minute {
			return errors.New("player can only participate in a challenge once per minute")
		}
	}

	challenge.ParticipationCount = lastChallenge.ParticipationCount + 1

	err = utils.ReadFromFile(config.AppConfig.ChallengesDataFilePath, &challenges)
	if err != nil {
		return err
	}

	challenges = append(challenges, *challenge)
	return utils.WriteToFile(config.AppConfig.ChallengesDataFilePath, challenges)
}

func (r *FileChallengeRepository) GetPendingChallenges() ([]models.Challenge, error) {
	var challenges []models.Challenge
	err := utils.ReadFromFile(config.AppConfig.ChallengesDataFilePath, &challenges)
	if err != nil {
		return nil, err
	}

	pendingChallenges := []models.Challenge{}
	for _, challenge := range challenges {
		if challenge.Status == models.Pending {
			pendingChallenges = append(pendingChallenges, challenge)
		}
	}
	return pendingChallenges, nil
}

func (r *FileChallengeRepository) UpdateStatus(challenge models.Challenge, status models.ChallengeStatus) error {
	var challenges []models.Challenge
	err := utils.ReadFromFile(config.AppConfig.ChallengesDataFilePath, &challenges)
	if err != nil {
		return err
	}

	for i, ch := range challenges {
		if ch.ID == challenge.ID && ch.PlayerID == challenge.PlayerID {
			challenges[i].Status = status
			break
		}
	}
	return utils.WriteToFile(config.AppConfig.ChallengesDataFilePath, challenges)
}
