package routes

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetInstallment(c *gin.Context) {
	GetEntities(c, func() *Installment {
		return &Installment{}
	})
}

func AddInstallment(c *gin.Context) {
	AddEntity(c, func() *Installment {
		return &Installment{}
	})
}

func UpdateInstallment(c *gin.Context) {
	UpdateEntity(c, func() *Installment {
		return &Installment{}
	})
}

func RemoveInstallment(c *gin.Context) {
	RemoveEntity(c, func() *Installment {
		return &Installment{}
	})
}

func (i *Installment) Create() (int, error) {
	if i.StartDate == "" {
		i.StartDate = time.Now().Format("2006-01-02 15:04:05")
	}
	if i.PayDate == "" {
		i.PayDate = time.Now().Format("2006-01-02 15:04:05")
	}

	if i.InstallmentsAmount == 0 {
		i.InstallmentsAmount = i.TotalAmount / float64(i.TotalInstallments)
	}

	id, err := i.insert()
	if err != nil {
		return 0, fmt.Errorf("error creating installment: %v", err)
	}

	for j := 1; j <= i.TotalInstallments; j++ {
		t := Transaction{
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
		}
		_, err := t.Create()

		if err != nil {
			i.DeleteEntity(id)
			return 0, fmt.Errorf("Error creating transaction %d: %v", j+1, err)
		}
	}

	return int(id), nil
}

func (i *Installment) insert() (int, error) {
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

func (i Installment) SetEntity(id int) (int, error) {
	query := `
    UPDATE InstallmentPlans 
    SET title = ?,
    total_amount = ?, 
    total_installments = ?,
    installment_amount = ?,
    start_date = ?, 
    pay_date = ?,
    status = ?,
    category_id = ?,
    subcategory_id = ? 
    WHERE id = ?`

	res, err := db.DB.Exec(query,
		i.Title, i.TotalAmount, i.TotalInstallments, i.InstallmentsAmount,
		i.StartDate, i.PayDate, i.Status, i.CategoryID, i.SubCategoryID, id,
	)
	if err != nil {
		return 0, err
	}

	idR, err := res.RowsAffected()

	return int(idR), err
}

func (i *Installment) DeleteEntity(id int) (int, error) {
	// Delete all transactions related to the installment
	_, err := DeleteInstallmentsTransactions(id)
	if err != nil {
		return 0, err
	}

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

func DeleteInstallmentsTransactions(id int) (int, error) {
	query := "SELECT * FROM Transactions WHERE installment_plan_id = ?"

	rows, err := db.DB.Query(query, id)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Transaction
		if err := t.Scan(rows); err != nil {
			return 0, err
		}
		_, err := t.DeleteEntity(id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (Installment) TableName() string {
	return "InstallmentPlans"
}

func (i Installment) GetSelectQuery() string {
	return fmt.Sprintf(`
        SELECT * 
        FROM %s
        ORDER BY ID DESC 
        LIMIT ? OFFSET ?`, i.TableName())
}

func (i *Installment) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&i.ID, &i.Title, &i.TotalAmount, &i.TotalInstallments,
		&i.InstallmentsAmount, &i.StartDate, &i.PayDate, &i.Status,
		&i.CategoryID, &i.SubCategoryID,
	)
}
