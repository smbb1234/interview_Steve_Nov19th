package tests

import (
	"steveInterviewMod/config"
	"steveInterviewMod/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

func SetupTestPath() {
	config.AppConfig.PlayerDataFilePath = "../testdata/players.json"
	config.AppConfig.LevelDataFilePath = "../testdata/levels.json"
	config.AppConfig.RoomDataFilePath = "../testdata/rooms.json"
	config.AppConfig.ReservationDataFilePath = "../testdata/reservation.json"
}
