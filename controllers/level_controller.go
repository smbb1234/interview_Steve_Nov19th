package controllers

import (
	"net/http"
	"steveInterviewMod/services"
	"steveInterviewMod/utils"

	"github.com/gin-gonic/gin"
)

func GetLevels(c *gin.Context) {
	levels, err := services.GetAllLevels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(levels))
}

func CreateLevel(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid input"))
		return
	}
	level, err := services.CreateLevel(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseSuccess(level))
}
