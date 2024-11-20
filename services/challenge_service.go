package services

import (
	"errors"
	"math"
	"math/rand/v2"
	"steveInterviewMod/models"
	"steveInterviewMod/repositories"
	"time"
)

var challengeID uint = 12

var challengeRepo repositories.ChallengeRepository = repositories.UniChallengeRepository()

func GetRecentChallenges() ([]models.Challenge, error) {
	return challengeRepo.GetRecentChallenges()
}

func ParticipateChallenge(request map[string]interface{}) (*models.Challenge, error) {

	playerId, ok := request["player_id"].(float64)
	if !ok {
		return nil, errors.New("player information is required")
	}

	amount, ok := request["amount"].(float64)
	if !ok {
		return nil, errors.New("amount is required")
	}
	epsilon := 1e-6 // Use absolute error comparison method to avoid floating point precision problems
	if math.Abs(amount-20.01) > epsilon {
		return nil, errors.New("invalid amount. Amount must be 20.01")
	}

	challenge := models.Challenge{
		ID:                 challengeID,
		PlayerID:           uint(playerId),
		ParticipationCount: 1,
		Status:             models.Pending,
		Timestamp:          time.Now().Format("2006-01-02 15:04:05"),
	}

	err := challengeRepo.Create(&challenge)
	if err != nil {
		return nil, err
	}
	return &challenge, nil
}

func UpdateChallengeIDPeriodically() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		UpdateChallengeID()
		UpdateChallengeStatus()
	}
}

func UpdateChallengeID() {
	challengeID += 1
}

func UpdateChallengeStatus() {
	pendingchallenges, err := challengeRepo.GetPendingChallenges()
	if err != nil {
		return
	}
	for _, challenge := range pendingchallenges {
		if rand.Float64() <= 0.01 {
			_ = challengeRepo.UpdateStatus(challenge, models.Win)
		} else {
			_ = challengeRepo.UpdateStatus(challenge, models.Lose)
		}

	}
}
