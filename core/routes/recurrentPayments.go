package routes

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func GetRecurringPayments(c *gin.Context) {
	GetEntities(c, func() *models.RecurringPayment {
		return &models.RecurringPayment{}
	})
}

func AddRecurringPayment(c *gin.Context) {
	AddEntity(c, func() *models.RecurringPayment {
		return &models.RecurringPayment{}
	})
}

func UpdateRecurringPayment(c *gin.Context) {
	UpdateEntity(c, func() *models.RecurringPayment {
		return &models.RecurringPayment{}
	})
}

func RemoveRecurringPayment(c *gin.Context) {
	RemoveEntity(c, func() *models.RecurringPayment {
		return &models.RecurringPayment{}
	})
}
