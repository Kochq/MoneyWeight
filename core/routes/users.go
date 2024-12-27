package routes

import (
	"api/db"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type transactionBody struct {
	Id                 int     `json:"id"`
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

func GetTransactions(c *gin.Context) {

	rows, err := db.DB.Query("SELECT * FROM Transactions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []transactionBody
	for rows.Next() {
		var t transactionBody
		err := rows.Scan(
			&t.Id,
			&t.Title,
			&t.Amount,
			&t.Category_id,
			&t.SubCategory_id,
			&t.Currency,
			&t.Payment_method,
			&t.Exchange_rate,
			&t.Notes,
			&t.Date,
			&t.InstallmentPlanID,
			&t.RecurringPaymentID,
			&t.Payment_number,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		transactions = append(transactions, t)
	}

	c.JSON(http.StatusOK, transactions)
}

func AddTransaction(c *gin.Context) {
	var body transactionBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid request",
		})
	}

	if body.Date == "" {
		body.Date = time.Now().Format("2006-01-02 15:04:05")
	}

	query := `
        INSERT INTO Transactions (title, amount, category_id, subcategory_id, currency, payment_method, exchange_rate, notes, date, installment_plan_id, recurring_payment_id, payment_number)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := db.DB.Exec(query,
		body.Title,
		body.Amount,
		body.Category_id,
		body.SubCategory_id,
		body.Currency,
		body.Payment_method,
		body.Exchange_rate,
		body.Notes,
		body.Date,
		body.InstallmentPlanID,
		body.RecurringPaymentID,
		body.Payment_number,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Transaction created with ID: %d", id))
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	query := "DELETE FROM Transactions WHERE id = ?"
	res, err := db.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID Not found"})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Transaction with ID: %s deleted", id))
}
