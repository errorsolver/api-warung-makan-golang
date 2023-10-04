package controllers

import (
	"net/http"

	"golang-wm-api/models"

	"github.com/gin-gonic/gin"
)

func (ControllerCollection) CreateTransactionDetails(c *gin.Context) {
	tx := models.DB.Begin()

	var transactionDetail models.TransactionDetail
	var product models.Product

	// if err := c.ShouldBindJSON(&transactionDetail); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := c.ShouldBindJSON(&transactionDetail); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "JSON Bind Error TransactionDetail",
			"error":   err.Error(),
		})
		return
	}

	// if err := models.DB.Create(&transactionDetail).Error; err != nil {
	// 	// tx.Rollback()
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error td": err.Error()})
	// 	return
	// }

	if err := tx.Create(&models.TransactionDetail{
		UserID:    transactionDetail.UserID,
		ProductID: transactionDetail.ProductID,
		Amount:    transactionDetail.Amount,
	}).Error; err != nil {
		// tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error td": err.Error()})
		return
	}

	// if err := models.DB.Create(&models.Transaction{
	// 	TotalPrice:          transaction.TotalPrice,
	// 	TransactionDetailID: transaction.TransactionDetailID,
	// }).Error; err != nil {
	// 	// tx.Rollback()
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error t": err.Error()})
	// 	return
	// }
	// if err := models.DB.Create(&transaction).Error; err != nil {
	// 	// tx.Rollback()
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error t": err.Error()})
	// 	return
	// }
	// tx.Commit()

	if err := tx.First(&product).Error; err != nil {
		tx.Rollback()
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Fail to get product data",
			"error":   err.Error(),
		})
		return
	}

	if err := tx.Last(&transactionDetail).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Fail to get product data",
			"error":   err.Error(),
		})
		return
	}

	totalPrice := int32(product.Price * transactionDetail.Amount)
	transaction := models.Transaction{
		TotalPrice:          totalPrice,
		TransactionDetailID: uint8(transactionDetail.ID),
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Fail to create transaction",
			"error":   err.Error(),
		})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction success",
		"data":    transactionDetail})
}

// func (ControllerCollection) GetTransactionDetails(c *gin.Context) {
// 	var transactions []models.Transaction

// 	models.DB.Find(&transactions)
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Success get all transactions",
// 		"data":    transactions,
// 	})
// }

func (ControllerCollection) GetAllTransaction(c *gin.Context) {
	var transactionDetails []models.TransactionDetail

	if err := models.DB.Preload("Transaction").Find(&transactionDetails).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Fail to get transaction detail",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get all transaction",
		"data":    transactionDetails,
	})
}

func (ControllerCollection) GetTransactionByID(c *gin.Context) {
	var transactionDetail models.TransactionDetail
	id := c.Param("id")

	if err := models.DB.Preload("Transaction").Find(&transactionDetail, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": "Fail to find transaction data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success get transaction from ID:",
		"data":    transactionDetail,
	})
}

// func (ControllerCollection) UpdateTransaction(c *gin.Context) {
// 	var transactionDetail models.TransactionDetail
// 	var transaction models.Transaction

// 	id:=c.Param("id")
// 	id:=c.BindJSON()

// 	if err:=c.ShouldBindJSON(transactionDetail).Error; err != nil{

// 	}
// }
