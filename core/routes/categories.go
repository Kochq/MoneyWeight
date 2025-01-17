package routes

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	GetEntities(c, func() *models.Categories {
		return &models.Categories{}
	})
}

func AddCategory(c *gin.Context) {
	AddEntity(c, func() *models.Categories {
		return &models.Categories{}
	})
}

func UpdateCategory(c *gin.Context) {
	UpdateEntity(c, func() *models.Categories {
		return &models.Categories{}
	})
}

func RemoveCategory(c *gin.Context) {
	RemoveEntity(c, func() *models.Categories {
		return &models.Categories{}
	})
}
