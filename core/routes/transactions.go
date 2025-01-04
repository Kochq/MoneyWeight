package routes

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	GetEntities(c, func() *models.Transaction {
		return &models.Transaction{}
	})
}

func AddTransaction(c *gin.Context) {
	AddEntity(c, func() *models.Transaction {
		return &models.Transaction{}
	})
}

func UpdateTransaction(c *gin.Context) {
	UpdateEntity(c, func() *models.Transaction {
		return &models.Transaction{}
	})
}

func RemoveTransaction(c *gin.Context) {
	RemoveEntity(c, func() *models.Transaction {
		return &models.Transaction{}
	})
}
