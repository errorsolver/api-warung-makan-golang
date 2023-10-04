package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-wm-api/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// type ProductControllers struct{}

func (ControllerCollection) GetProducts(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success get all products",
		"data":    products,
	})
}

func (ControllerCollection) GetProductById(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
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
		"message": "Success get product",
		"data":    product,
	})
}

func (ControllerCollection) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success create product",
		"data":    product,
	})
}

func (ControllerCollection) UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Fail update data"})
		return
	}

	if err := models.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Product output error",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update data",
		"data":    product,
	})
}

func (ControllerCollection) DeleteProduct(c *gin.Context) {
	var product models.Product

	// userInput := map[string]string{"id": "0"}
	var userInput struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := userInput.Id.Int64()
	if err := models.DB.Delete(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Fail delete data",
			"error":   err.Error(),
		})
		return
	}
	// if models.DB.Delete(&product, id).RowsAffected == 0 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Fail delete data"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("success delete user: %s", product.ProductName),
	})
}
