package routes

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func GetTransactionDetails(c *gin.Context) {
	GetEntities(c, func() *models.TransactionDetails {
		return &models.TransactionDetails{}
	})
}

func AddTransactionDetails(c *gin.Context) {
	AddEntity(c, func() *models.TransactionDetails {
		return &models.TransactionDetails{}
	})
}

func UpdateTransactionDetails(c *gin.Context) {
	UpdateEntity(c, func() *models.TransactionDetails {
		return &models.TransactionDetails{}
	})
}

func RemoveTransactionDetails(c *gin.Context) {
	RemoveEntity(c, func() *models.TransactionDetails {
		return &models.TransactionDetails{}
	})
}
