package controllers

import (
	"golang-wm-api/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// type UserControllers struct{}

func (ControllerCollection) GetUsers(c *gin.Context) {
	var user []models.User

	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Seccess get all users",
		"data":    user,
	})
}

func (ControllerCollection) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err := models.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success create user",
	})
}
