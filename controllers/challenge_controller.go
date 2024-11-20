package controllers

import (
	"net/http"
	"steveInterviewMod/services"
	"steveInterviewMod/utils"

	"github.com/gin-gonic/gin"
)

func ParticipateChallenge(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid input"))
		return
	}
	challenge, err := services.ParticipateChallenge(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseSuccess(challenge))
}

func GetChallengeResults(c *gin.Context) {
	results, err := services.GetRecentChallenges()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(results))
}
