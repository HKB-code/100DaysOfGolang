package routes

import (
	"net/http"
	"project/REST-API/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request."})
		return
	}
	userId := context.GetInt64("userId")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event."})
		return
	}
	err = event.RegisterEvent(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered."})
}

func cancelRegister(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request."})
		return
	}
	userId := context.GetInt64("userId")
	var event models.Event
	event.Id = eventId
	err = event.DeleteRegister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registeration Delete."})
}
