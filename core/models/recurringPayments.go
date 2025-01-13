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
	Status      bool    `json:"status"`
	StartDate     string  `json:"start_date"`
	PayDate       string  `json:"pay_date"`
	Frequency     string  `json:"frequency"`
	FromAccountID int     `json:"from_account_id"`
	Payments      int     `json:"payments"`
}

func (rp *RecurringPayment) Create() (int, error) {
	if rp.StartDate == "" {
		rp.StartDate = time.Now().Format("2006-01-02 15:04:05")
	}
	if rp.PayDate == "" {
		rp.PayDate = time.Now().Format("2006-01-02 15:04:05")
	}

	if rp.Frequency == "" {
		rp.Frequency = "monthly"
	}

	return rp.Insert()
}

func (rp *RecurringPayment) Insert() (int, error) {
	query := `
    INSERT INTO RecurringPayments (
    title, amount, category_id, subcategory_id, status, start_date,
    pay_date, frequency, from_account_id, payments
    )
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := db.DB.Exec(query,
		&rp.Title, &rp.Amount, &rp.CategoryID, &rp.SubCategoryID, &rp.Status,
		&rp.StartDate, &rp.PayDate, &rp.Frequency, &rp.FromAccountID, &rp.Payments,
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
    status = ?,
    start_date = ?,
    pay_date = ?,
    frequency = ?
    from_account_id = ?
    payments = ?
    WHERE id = ?`

	res, err := db.DB.Exec(query,
		&rp.Title, &rp.Amount, &rp.CategoryID, &rp.SubCategoryID, &rp.Status,
		&rp.StartDate, &rp.PayDate, &rp.Frequency, &rp.FromAccountID,
		&rp.Payments, id,
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
		&rp.ID, &rp.Title, &rp.Amount, &rp.Frequency, &rp.StartDate,
		&rp.PayDate, &rp.Status, &rp.CategoryID, &rp.SubCategoryID,
		&rp.FromAccountID, &rp.Payments,
	)
}
