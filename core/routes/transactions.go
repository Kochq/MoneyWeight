package routes

import (
	"api/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	// Agregar paginación
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	query := `
    SELECT * 
    FROM Transactions 
    ORDER BY date DESC 
    LIMIT ? OFFSET ?`

	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	transactions := []Transaction{}
	for rows.Next() {
		var t Transaction
		if err := scanTransaction(rows, &t); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
			return
		}
		transactions = append(transactions, t)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   transactions,
	})
}

func AddTransaction(c *gin.Context) {
	var body Transaction

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request body",
		})
		return
	}

	if body.Date == "" {
		body.Date = time.Now().Format("2006-01-02 15:04:05")
	}

	id, err := insertTransaction(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"id":     id,
	})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM Transactions WHERE id = ?"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Database error",
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Could not get affected rows",
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  "Transaction not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Transaction with ID %s deleted", id),
	})
}

func UpdateTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Invalid ID",
		})
	}

	var body Transaction

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request body",
		})
		return
	}

	rowsAffected, err := updateTransaction(body, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// It's working perfectly fine...
	// it wont affect any rows if it can't change anything
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  "Transaction not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Transaction with ID %d updated", id),
	})
}

func insertTransaction(t Transaction) (int64, error) {
	query := `
    INSERT INTO Transactions (
    title, amount, category_id, subcategory_id, 
    currency, payment_method, exchange_rate, 
    notes, date, installment_plan_id, 
    recurring_payment_id, payment_number
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query,
		t.Title, t.Amount, t.CategoryID, t.SubCategoryID,
		t.Currency, t.PaymentMethod, t.ExchangeRate,
		t.Notes, t.Date, t.InstallmentPlanID,
		t.RecurringPaymentID, t.PaymentNumber)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func scanTransaction(rows *sql.Rows, t *Transaction) error {
	return rows.Scan(
		&t.ID, &t.Title, &t.Amount, &t.CategoryID,
		&t.SubCategoryID, &t.Currency, &t.PaymentMethod,
		&t.ExchangeRate, &t.Notes, &t.Date,
		&t.InstallmentPlanID, &t.RecurringPaymentID,
		&t.PaymentNumber,
	)
}

func updateTransaction(t Transaction, id int) (int64, error) {
	query := `
    UPDATE Transactions 
    SET title = ?, 
    amount = ?, 
    category_id = ?, 
    subcategory_id = ?,
    currency = ?, 
    payment_method = ?,
    exchange_rate = ?,
    notes = ?,
    date = ?,
    installment_plan_id = ?,
    recurring_payment_id = ?,
    payment_number = ?
    WHERE id = ?
    `

	res, err := db.DB.Exec(query,
		t.Title, t.Amount, t.CategoryID, t.SubCategoryID,
		t.Currency, t.PaymentMethod, t.ExchangeRate,
		t.Notes, t.Date, t.InstallmentPlanID,
		t.RecurringPaymentID, t.PaymentNumber, id,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}