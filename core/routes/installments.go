package routes

import (
	"api/db"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetInstallment(c *gin.Context) {
	query := `SELECT * FROM InstallmentPlans`

	rows, err := db.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	installments := []Installment{}
	for rows.Next() {
		var i Installment
		if err := scanInstallment(rows, &i); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		installments = append(installments, i)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   installments,
	})
}

func AddInstallment(c *gin.Context) {
	var body Installment
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request body",
		})
		return
	}

	id, err := CreateInstallment(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"id":     id,
	})
}

func CreateInstallment(i Installment) (int, error) {
	if i.StartDate == "" {
		i.StartDate = time.Now().Format("2006-01-02 15:04:05")
	}
	if i.PayDate == "" {
		i.PayDate = time.Now().Format("2006-01-02 15:04:05")
	}

	if i.InstallmentsAmount == 0 {
		i.InstallmentsAmount = i.TotalAmount / float64(i.TotalInstallments)
	}

	id, err := insertInstallment(i)
	if err != nil {
		return 0, fmt.Errorf("error creating installment: %v", err)
	}

	createdTransactions := []int{}
	for j := 1; j <= i.TotalInstallments; j++ {
		transactionId, err := CreateTransaction(Transaction{
			Title:             i.Title,
			Amount:            i.InstallmentsAmount,
			CategoryID:        i.CategoryID,
			SubCategoryID:     i.SubCategoryID,
			Currency:          "USD",
			PaymentMethod:     "Cash",
			ExchangeRate:      1,
			Notes:             "",
			Date:              i.PayDate,
			InstallmentPlanID: &id,
			PaymentNumber:     &j,
		})

		if err != nil {
			for _, tId := range createdTransactions {
				DeleteTransaction(tId)
			}

			DeleteInstallment(id)
			return 0, fmt.Errorf("Error creating transaction %d: %v", j+1, err)
		}

		createdTransactions = append(createdTransactions, transactionId)
	}

	return int(id), nil
}

func DeleteInstallment(id int) (int, error) {
	query := "DELETE FROM InstallmentPlans WHERE id = ?"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("Error deleting installment %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error trying to retrieve rows affected")
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("Couldnt remove installment %d", id)
	}
	return id, nil
}

func insertInstallment(i Installment) (int, error) {
	query := `
    INSERT INTO InstallmentPlans (
    title, total_amount, total_installments, installment_amount, 
    start_date, pay_date, status, category_id, subcategory_id
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query,
		i.Title, i.TotalAmount, i.TotalInstallments, i.InstallmentsAmount,
		i.StartDate, i.PayDate, i.Status, i.CategoryID, i.SubCategoryID,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func scanInstallment(rows *sql.Rows, i *Installment) error {
	return rows.Scan(
		&i.ID, &i.Title, &i.TotalAmount, &i.TotalInstallments, &i.InstallmentsAmount,
		&i.StartDate, &i.PayDate, &i.Status, &i.CategoryID, &i.SubCategoryID,
	)
}