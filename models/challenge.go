package models

type ChallengeStatus string

const (
	Win     ChallengeStatus = "Win"
	Lose    ChallengeStatus = "Lose"
	Pending ChallengeStatus = "Pending"
)

type Challenge struct {
	ID                 uint            `json:"id"`
	PlayerID           uint            `json:"player_id"`
	ParticipationCount int             `json:"participation_count"`
	Status             ChallengeStatus `json:"status"`
	Timestamp          string          `json:"timestamp"`
}
