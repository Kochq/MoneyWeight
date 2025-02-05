package routes

import (
	"api/models"
	"github.com/gin-gonic/gin"
)

func GetSubCategories(c *gin.Context) {
	GetEntities(c, func() *models.SubCategories {
		return &models.SubCategories{}
	})
}

func AddSubCategory(c *gin.Context) {
	AddEntity(c, func() *models.SubCategories {
		return &models.SubCategories{}
	})
}

func UpdateSubCategory(c *gin.Context) {
	UpdateEntity(c, func() *models.SubCategories {
		return &models.SubCategories{}
	})
}

func RemoveSubCategory(c *gin.Context) {
	RemoveEntity(c, func() *models.SubCategories {
		return &models.SubCategories{}
	})
}
