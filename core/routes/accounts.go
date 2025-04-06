package routes

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func GetAccounts(c *gin.Context) {
	GetEntities(c, func() *models.Account {
		return &models.Account{}
	})
}

func AddAccount(c *gin.Context) {
	AddEntity(c, func() *models.Account {
		return &models.Account{}
	})
}

func UpdateAccount(c *gin.Context) {
	UpdateEntity(c, func() *models.Account {
		return &models.Account{}
	})
}

func RemoveAccount(c *gin.Context) {
	RemoveEntity(c, func() *models.Account {
		return &models.Account{}
	})
}
