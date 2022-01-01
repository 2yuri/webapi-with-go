package controllers

import (
	"github.com/LOO2/business-remote-management-api/database"
	"github.com/LOO2/business-remote-management-api/models"
	"github.com/gin-gonic/gin"
)

func ShowAllRevenues(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Revenue
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find revenue: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateRevenue(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Revenue

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create revenue: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
