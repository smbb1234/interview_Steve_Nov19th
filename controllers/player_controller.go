package controllers

import (
	"net/http"
	"steveInterviewMod/services"
	"steveInterviewMod/utils"

	"github.com/gin-gonic/gin"
)

func GetPlayers(c *gin.Context) {
	players, err := services.GetAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(players))
}

func CreatePlayer(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid input"))
		return
	}
	player, err := services.CreatePlayer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseSuccess(player))
}

func GetPlayerByID(c *gin.Context) {
	id := c.Param("id")
	player, err := services.GetPlayerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(player))
}

func UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid input"))
		return
	}
	player, err := services.UpdatePlayer(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(player))
}

func DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	err := services.DeletePlayer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess("Player deleted successfully"))
}
