package main

import (
	"api/db"
	"api/routes"

	"github.com/gin-gonic/gin"
)

type postRequestBody struct {
	Title              string  `json:"title"`
	Amount             float64 `json:"amount"`
	Category_id        int     `json:"category_id"`
	SubCategory_id     int     `json:"subcategory_id"`
	Currency           string  `json:"currency"`
	Payment_method     string  `json:"payment_method"`
	Exchange_rate      float64 `json:"exchange_rate"`
	Notes              string  `json:"notes"`
	Date               string  `json:"date"`
	InstallmentPlanID  *int    `json:"installment_plan_id"`
	RecurringPaymentID *int    `json:"recurring_payment_id"`
	Payment_number     *int    `json:"payment_number"`
}

func main() {
	db.ConnectDB()
	defer db.DB.Close()

	r := gin.Default()

	r.GET("/users", routes.GetTransactions)

	r.POST("/post", routes.AddTransaction)
	r.DELETE("/delete/:id", routes.DeleteTransaction)

	r.Run(":8080") // Inicia el servidor en el puerto 8080
}
