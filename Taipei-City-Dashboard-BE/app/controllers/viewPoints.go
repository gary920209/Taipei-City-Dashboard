package controllers

import (
	"net/http"
	"strconv"

	"TaipeiCityDashboardBE/app/models"

	"github.com/gin-gonic/gin"
)

/*
CreateViewPoint creates a new viewpoint for the user
POST /api/v1/user/:id/viewpoint
*/
func CreateViewPoint(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var viewPoint models.ViewPoints
	if err := c.ShouldBindJSON(&viewPoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	viewPoint, err = models.CreateViewPoint(userId, viewPoint.CenterX, viewPoint.CenterY, viewPoint.Zoom, viewPoint.Pitch, viewPoint.Bearing, viewPoint.Name, viewPoint.PointType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Viewpoint created", "data": viewPoint})
}

/*
GetViewPointByUserID retrieves the viewpoint for the user
GET /api/v1/user/:id/viewpoint
*/
func GetViewPointByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	viewPoint, err := models.GetViewPointByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, viewPoint)
}

/*
DeleteViewPoint deletes the viewpoint
DELETE /api/v1/user/:id/viewpoint
*/
func DeleteViewPoint(c *gin.Context) {
	pointId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid viewpoint ID"})
		return
	}

	err = models.DeleteViewPoint(pointId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Viewpoint deleted"})
}
