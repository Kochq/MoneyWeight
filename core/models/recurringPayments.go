package models

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"
)

type RecurringPayment struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Amount        float64 `json:"amount"`
	CategoryID    int     `json:"category_id"`
	SubCategoryID int     `json:"subcategory_id"`
	IsActive      bool    `json:"is_active"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	Frequency     string  `json:"frequency"`
}

func (rp *RecurringPayment) Create() (int, error) {
	if rp.StartDate == "" {
		rp.StartDate = time.Now().Format("2006-01-02 15:04:05")
	}
	if rp.EndDate == "" {
		rp.EndDate = time.Now().Format("2006-01-02 15:04:05")
	}

	if rp.Frequency == "" {
		rp.Frequency = "monthly"
	}

	return rp.Insert()
}

func (rp *RecurringPayment) Insert() (int, error) {
	query := `
    INSERT INTO RecurringPayments (
    title, amount, category_id, subcategory_id, is_active, start_date,
    end_date, frequency
    )
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := db.DB.Exec(query,
		&rp.Title, &rp.Amount, &rp.CategoryID, &rp.SubCategoryID, &rp.IsActive,
		&rp.StartDate, &rp.EndDate, &rp.Frequency,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func (rp *RecurringPayment) SetEntity(id int) (int, error) {
	query := `
    UPDATE RecurringPayments 
    SET title = ?, 
    amount = ?,
    category_id = ?,
    subcategory_id = ?,
    is_active = ?,
    start_date = ?,
    end_date = ?,
    frequency = ?
    WHERE id = ?`

	res, err := db.DB.Exec(query,
		&rp.Title, &rp.Amount, &rp.CategoryID, &rp.SubCategoryID, &rp.IsActive,
		&rp.StartDate, &rp.EndDate, &rp.Frequency, id,
	)
	if err != nil {
		return 0, err
	}

	idR, err := res.RowsAffected()

	return int(idR), err
}

func (rp *RecurringPayment) DeleteEntity(id int) (int, error) {
	query := "DELETE FROM RecurringPayments WHERE id = ?"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("Error deleting transaction %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error trying to retrieve rows affected")
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("Couldnt remove recurring payment %d", id)
	}

	return id, nil
}

func (RecurringPayment) TableName() string {
	return "RecurringPayments"
}

func (rp RecurringPayment) GetSelectQuery() string {
	return fmt.Sprintf(`
    SELECT * 
    FROM %s
    ORDER BY ID DESC 
    LIMIT ? OFFSET ?`, rp.TableName())
}

func (rp *RecurringPayment) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&rp.ID, &rp.Title, &rp.Amount, &rp.Frequency, &rp.StartDate, &rp.EndDate,
		&rp.IsActive, &rp.CategoryID, &rp.SubCategoryID,
	)
}
