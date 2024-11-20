package config

type Config struct {
	PlayerDataFilePath      string
	LevelDataFilePath       string
	RoomDataFilePath        string
	ReservationDataFilePath string
	LogDataFilePath         string
	ChallengesDataFilePath  string
	UseSQLDatabase          bool
}

var AppConfig Config

func Setup() {

	AppConfig = Config{
		PlayerDataFilePath:      "testdata/players.json",
		LevelDataFilePath:       "testdata/levels.json",
		RoomDataFilePath:        "testdata/rooms.json",
		ReservationDataFilePath: "testdata/reservations.json",
		LogDataFilePath:         "testdata/logs.json",
		ChallengesDataFilePath:  "testdata/challenges.json",
		UseSQLDatabase:          false,
	}
}
