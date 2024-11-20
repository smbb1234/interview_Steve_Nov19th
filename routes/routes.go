package routes

import (
	"steveInterviewMod/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Player Routes
	playerRoutes := r.Group("/players")
	{
		playerRoutes.GET("", controllers.GetPlayers)
		playerRoutes.POST("", controllers.CreatePlayer)
		playerRoutes.GET(":id", controllers.GetPlayerByID)
		playerRoutes.PUT(":id", controllers.UpdatePlayer)
		playerRoutes.DELETE(":id", controllers.DeletePlayer)
	}

	// Level Routes
	levelRoutes := r.Group("/levels")
	{
		levelRoutes.GET("", controllers.GetLevels)
		levelRoutes.POST("", controllers.CreateLevel)
	}

	// Room Routes
	roomRoutes := r.Group("/rooms")
	{
		roomRoutes.GET("", controllers.GetRooms)
		roomRoutes.POST("", controllers.CreateRoom)
		roomRoutes.GET(":id", controllers.GetRoomByID)
		roomRoutes.PUT(":id", controllers.UpdateRoom)
		roomRoutes.DELETE(":id", controllers.DeleteRoom)
	}

	// Reservation Routes
	reservationsRoutes := r.Group("/reservations")
	{
		reservationsRoutes.GET("", controllers.GetReservations) // Supports room_id, date, and limit as query parameters
		reservationsRoutes.POST("", controllers.CreateReservation)
	}

	//Log Routes
	logRoutes := r.Group("/logs")
	{
		logRoutes.GET("", controllers.GetLogs) // Supports player_id, action, start_time, end_time, and limit as query parameters
		logRoutes.POST("", controllers.CreateLog)
	}

	// Challenge Routes
	challengeRoutes := r.Group("/challenges")
	{
		challengeRoutes.POST("", controllers.ParticipateChallenge)
		challengeRoutes.GET("/results", controllers.GetChallengeResults)
	}

	// Payment Routes
	// paymentRoutes := r.Group("/payments")
	// {
	// 	paymentRoutes.POST("", controllers.PrecessPayment)
	// 	paymentRoutes.GET(":id", controllers.GetPaymentByID)
	// }
}
