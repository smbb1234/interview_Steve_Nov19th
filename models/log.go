package models

import (
	"errors"
)

type Action string

const (
	Register        Action = "register"
	Login           Action = "login"
	Logout          Action = "logout"
	EnterRoom       Action = "enter_room"
	LeaveRoom       Action = "leave_room"
	JoinChallenge   Action = "join_challenge"
	ChallengeResult Action = "challenge_result"
)

var actions = []Action{Register, Login, Logout, EnterRoom, LeaveRoom, JoinChallenge, ChallengeResult}

func ToAction(actionStr string) (Action, error) {
	for _, action := range actions {
		if string(action) == actionStr {
			return action, nil
		}
	}
	return "", errors.New("invalid action value: " + actionStr)
}

type Log struct {
	ID        uint   `json:"id"`
	PlayerID  uint   `json:"player_id"`
	Action    Action `json:"action"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}
