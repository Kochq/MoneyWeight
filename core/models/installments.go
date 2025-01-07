package models

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"
)

type Installment struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	TotalAmount        float64 `json:"total_amount"`
	TotalInstallments  int     `json:"total_installments"`
	InstallmentsAmount float64 `json:"installment_amount"`
	StartDate          string  `json:"start_date"` // Podríamos usar time.Time
	PayDate            string  `json:"pay_date"`   // Podríamos usar time.Time
	Status             string  `json:"status"`
	CategoryID         int     `json:"category_id"`
	SubCategoryID      int     `json:"subcategory_id"`
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

	id, err := i.Insert()
	if err != nil {
		return 0, fmt.Errorf("error creating installment: %v", err)
	}

	for j := 1; j <= i.TotalInstallments; j++ {
        fDate, err := time.Parse("2006-01-02 15:04:05", i.PayDate)
		if err != nil {
			return 0, fmt.Errorf("Error parsing start date: %v", err)
		}
		addedDate := fDate.AddDate(0, j, 0)
		fAddedDate := addedDate.Format("2006-01-02 15:04:05")

		t := Transaction{
			Title:             i.Title,
			Amount:            i.InstallmentsAmount,
			CategoryID:        i.CategoryID,
			SubCategoryID:     i.SubCategoryID,
			Currency:          "USD",
			PaymentMethod:     "Cash",
			ExchangeRate:      1,
			Notes:             "",
			Date:              fAddedDate,
			InstallmentPlanID: &id,
			PaymentNumber:     &j,
		}
		_, err = t.Create()

		if err != nil {
			i.DeleteEntity(id)
			return 0, fmt.Errorf("Error creating transaction %d: %v", j+1, err)
		}
	}

	return int(id), nil
}

func (i *Installment) Insert() (int, error) {
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

		_, err := t.DeleteEntity(t.ID)
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
