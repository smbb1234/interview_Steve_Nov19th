package controllers

import (
	"net/http"
	"steveInterviewMod/services"
	"steveInterviewMod/utils"

	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	rooms, err := services.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(rooms))
}

func CreateRoom(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid input"))
		return
	}
	room, err := services.CreateRoom(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseSuccess(room))
}

func GetRoomByID(c *gin.Context) {
	id := c.Param("id")
	room, err := services.GetRoomByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(room))
}

func UpdateRoom(c *gin.Context) {
	id := c.Param("id")
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError("Invalid input"))
		return
	}
	room, err := services.UpdateRoom(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess(room))
}

func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteRoom(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess("Room deleted successfully"))
}
