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

	r.Run(":8080") // Inicia el servidor en el puerto 8080
}
