package controllers

import (
	"net/http"
	"steveInterviewMod/services"
	"steveInterviewMod/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReservations(c *gin.Context) {
	roomIDStr := c.Query("room_id")
	date := c.Query("date")
	limitStr := c.Query("limit")

	var roomID *uint
	if roomIDStr != "" {
		id, err := strconv.ParseUint(roomIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid room ID"))
			return
		}
		roomID = new(uint)
		*roomID = uint(id)
	}

	var limit *int
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid limit"))
			return
		}
		limit = &l
	}

	reservations, err := services.GetReservations(roomID, date, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(reservations))
}

func CreateReservation(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid input"))
		return
	}
	reservation, err := services.CreateReservation(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseSuccess(reservation))
}
