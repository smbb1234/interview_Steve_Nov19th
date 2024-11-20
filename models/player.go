package models

type Player struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	LevelID uint   `json:"level_id"`
	Others  string `json:"others"`
}
