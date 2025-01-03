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

func (Transaction) TableName() string {
	return "Transactions"
}

func (t *Transaction) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&t.ID, &t.Title, &t.Amount, &t.CategoryID, &t.SubCategoryID,
		&t.Currency, &t.PaymentMethod, &t.ExchangeRate, &t.Notes, &t.Date,
		&t.InstallmentPlanID, &t.RecurringPaymentID, &t.PaymentNumber,
	)
}

func (t Transaction) GetQuery() string {
	return fmt.Sprintf(`
        SELECT * 
        FROM %s 
        ORDER BY date DESC 
        LIMIT ? OFFSET ?`, t.TableName())

}

func (t *Transaction) Create() (int, error) {
	if t.Date == "" {
		t.Date = time.Now().Format("2006-01-02 15:04:05")
	}
	return t.insert()
}

func (t *Transaction) insert() (int, error) {
	query := `
    INSERT INTO Transactions (
    title, amount, category_id, subcategory_id,
    currency, payment_method, exchange_rate,
    notes, date, installment_plan_id,
    recurring_payment_id, payment_number
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query,
		t.Title, t.Amount, t.CategoryID, t.SubCategoryID, t.Currency,
		t.PaymentMethod, t.ExchangeRate, t.Notes, t.Date, t.InstallmentPlanID,
		t.RecurringPaymentID, t.PaymentNumber)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func GetTransactions(c *gin.Context) {
	GetEntities(c, func() *Transaction {
		return &Transaction{}
	})
}

func AddTransaction(c *gin.Context) {
	AddEntity(c, func() *Transaction {
		return &Transaction{}
	})
}

func RemoveTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Invalid ID",
		})
	}

	id, err = DeleteTransaction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Transaction with ID %d deleted", id),
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

func DeleteTransaction(id int) (int, error) {
	query := "DELETE FROM Transactions WHERE id = ?"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("Error deleting transaction %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error trying to retrieve rows affected")
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("Couldnt remove transaction %d", id)
	}

	return id, nil
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
		t.Title, t.Amount, t.CategoryID, t.SubCategoryID, t.Currency,
		t.PaymentMethod, t.ExchangeRate, t.Notes, t.Date, t.InstallmentPlanID,
		t.RecurringPaymentID, t.PaymentNumber, id,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
