package routes

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func GetInstallment(c *gin.Context) {
	GetEntities(c, func() *models.Installment {
		return &models.Installment{}
	})
}

func AddInstallment(c *gin.Context) {
	AddEntity(c, func() *models.Installment {
		return &models.Installment{}
	})
}

func UpdateInstallment(c *gin.Context) {
	UpdateEntity(c, func() *models.Installment {
		return &models.Installment{}
	})
}

func RemoveInstallment(c *gin.Context) {
	RemoveEntity(c, func() *models.Installment {
		return &models.Installment{}
	})
}
