package controllers

import (
	"net/http"
	"steveInterviewMod/services"
	"steveInterviewMod/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLogs(c *gin.Context) {
	playerIDStr := c.Query("player_id")
	action := c.Query("action")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	limitStr := c.Query("limit")

	var playerID *uint
	if playerIDStr != "" {
		id, err := strconv.ParseUint(playerIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid player ID"))
			return
		}
		playerID = new(uint)
		*playerID = uint(id)
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

	logs, err := services.GetLogsByCondition(playerID, action, startTime, endTime, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(logs))
}

func CreateLog(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid input"))
		return
	}
	log, err := services.CreateLog(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseSuccess(log))
}
