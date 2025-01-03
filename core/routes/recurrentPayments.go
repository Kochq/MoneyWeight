package routes

import (
	"api/db"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRecurringPayments(c *gin.Context) {
	query := `SELECT * FROM RecurringPayments`
	rows, err := db.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	recurringPayments := []RecurringPayment{}
	for rows.Next() {
		var rp RecurringPayment
		if err := scanRecurringPayment(rows, &rp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		recurringPayments = append(recurringPayments, rp)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   recurringPayments,
	})
}

func AddRecurringPayment(c *gin.Context) {
	var body RecurringPayment
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request body",
		})
		return
	}
	id, err := CreateRecurringPayment(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"id":     id,
	})
}

func scanRecurringPayment(rows *sql.Rows, rp *RecurringPayment) error {
	return rows.Scan(
		&rp.ID, &rp.Title, &rp.Amount, &rp.Frequency, &rp.StartDate, &rp.EndDate,
		&rp.IsActive, &rp.CategoryID, &rp.SubCategoryID,
	)
}

func CreateRecurringPayment(r RecurringPayment) (int, error) {
	if r.StartDate == "" {
		r.StartDate = time.Now().Format("2006-01-02 15:04:05")
	}
	if r.EndDate == "" {
		r.EndDate = time.Now().Format("2006-01-02 15:04:05")
	}

	if r.Frequency == "" {
		r.Frequency = "monthly"
	}

	return insertRecurringPayment(r)
}

func insertRecurringPayment(r RecurringPayment) (int, error) {
	query := `
    INSERT INTO RecurringPayments (
    title, amount, category_id, subcategory_id, is_active, start_date,
    end_date, frequency
    )
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := db.DB.Exec(query,
		&r.Title, &r.Amount, &r.CategoryID, &r.SubCategoryID, &r.IsActive,
		&r.StartDate, &r.EndDate, &r.Frequency,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}
