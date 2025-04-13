package main

import (
	"api/db"
	"api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectDB()
	defer db.DB.Close()

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/transactions/details", routes.GetTransactionDetails)
		api.GET("/transactions", routes.GetTransactions)
		api.POST("/transactions", routes.AddTransaction)
		api.PUT("/transactions/:id", routes.UpdateTransaction)
		api.DELETE("/transactions/:id", routes.RemoveTransaction)
	}

	{
		api.GET("/installments", routes.GetInstallment)
		api.POST("/installments", routes.AddInstallment)
		api.PUT("/installments/:id", routes.UpdateInstallment)
		api.DELETE("/installments/:id", routes.RemoveInstallment)
	}

	{
		api.GET("/recurring", routes.GetRecurringPayments)
		api.POST("/recurring", routes.AddRecurringPayment)
		api.PUT("/recurring/:id", routes.UpdateRecurringPayment)
		api.DELETE("/recurring/:id", routes.RemoveRecurringPayment)
	}

	{
		api.GET("/categories", routes.GetCategories)
		api.POST("/categories", routes.AddCategory)
		api.PUT("/categories/:id", routes.UpdateCategory)
		api.DELETE("/categories/:id", routes.RemoveCategory)
	}

	{
		api.GET("/subCategories", routes.GetSubCategories)
		api.POST("/subCategories", routes.AddSubCategory)
		api.PUT("/subCategories/:id", routes.UpdateSubCategory)
		api.DELETE("/subCategories/:id", routes.RemoveSubCategory)
	}

	{
		api.GET("/accounts", routes.GetAccounts)
		api.POST("/accounts", routes.AddAccount)
		api.PUT("/accounts/:id", routes.UpdateAccount)
		api.DELETE("/accounts/:id", routes.RemoveAccount)
	}

	r.Run(":8080") // Inicia el servidor en el puerto 8080
}
