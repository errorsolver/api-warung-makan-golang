package controllers

import (
	"encoding/json"
	"fmt"
	"golang-wm-api/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (ControllerCollection) GetUsers(c *gin.Context) {
	var user []models.User

	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Seccess get all users",
		"data":    user,
	})
}

func (ControllerCollection) GetUsersById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := models.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Server error",
				"error":   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get user " + id,
		"error":   user,
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

func (ControllerCollection) UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "Please input correct json data",
			"error":   err.Error(),
		})
	}
	if models.DB.Model(&user).Where("id", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Fail to update user data"})
		return
	}

	if err := models.DB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "User output error",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update user data",
		"data":    user,
	})
}

func (ControllerCollection) DeleteUser(c *gin.Context) {
	var user models.User

	var userInput struct {
		Id json.Number
	}
	id, _ := userInput.Id.Int64()

	if err := models.DB.Model(&user).Where("id = ?", id).Delete(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": "Fail to delete an user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("success delete user: %s", user.Username),
		"data":    user,
	})
}
