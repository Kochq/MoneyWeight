package models

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"
)

type Transaction struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Amount             float64 `json:"amount"`
	CategoryID         int     `json:"category_id"`
	SubCategoryID      int     `json:"subcategory_id"`
	Currency           string  `json:"currency"`
	PaymentMethod      string  `json:"payment_method"`
	ExchangeRate       float64 `json:"exchange_rate"`
	Notes              string  `json:"notes"`
	Date               string  `json:"date"` // Podríamos usar time.Time
	InstallmentPlanID  *int    `json:"installment_plan_id,omitempty"`
	RecurringPaymentID *int    `json:"recurring_payment_id,omitempty"`
	FromAccountID      int     `json:"from_account_id"`
	ToAccountID        *int64  `json:"to_account_id,omitempty"`
	PaymentNumber      *int    `json:"payment_number,omitempty"`
	Status             string  `json:"status"`
}

func (t *Transaction) Create() (int, error) {
	if t.Date == "" {
		t.Date = time.Now().Format("2006-01-02 15:04:05")
	}
	return t.Insert()
}

func (t *Transaction) Insert() (int, error) {
	query := `
    INSERT INTO Transactions (
    title, amount, category_id, subcategory_id, currency, payment_method,
    exchange_rate, notes, date, installment_plan_id, recurring_payment_id,
    payment_number, from_account_id, to_account_id, status
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query,
		t.Title, t.Amount, t.CategoryID, t.SubCategoryID, t.Currency,
		t.PaymentMethod, t.ExchangeRate, t.Notes, t.Date, t.InstallmentPlanID,
		t.RecurringPaymentID, t.PaymentNumber, t.FromAccountID, t.ToAccountID,
		t.Status)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (t Transaction) SetEntity(id int) (int, error) {
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
    from_account_id = ?
    to_account_id = ?
    status = ?
    WHERE id = ?
    `

	res, err := db.DB.Exec(query,
		t.Title, t.Amount, t.CategoryID, t.SubCategoryID, t.Currency,
		t.PaymentMethod, t.ExchangeRate, t.Notes, t.Date, t.InstallmentPlanID,
		t.RecurringPaymentID, t.PaymentNumber, t.FromAccountID, t.ToAccountID,
		t.Status, id,
	)
	if err != nil {
		return 0, err
	}

	idR, err := res.RowsAffected()

	return int(idR), err
}

func (t *Transaction) DeleteEntity(id int) (int, error) {
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

func (Transaction) TableName() string {
	return "Transactions"
}

func (t Transaction) GetSelectQuery() string {
	return fmt.Sprintf(`
        SELECT *
        FROM %s
        ORDER BY date DESC
        LIMIT ? OFFSET ?`, t.TableName())
}

func (t *Transaction) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&t.ID, &t.Title, &t.Amount, &t.CategoryID, &t.FromAccountID,
		&t.ToAccountID, &t.SubCategoryID, &t.Currency, &t.PaymentMethod,
		&t.ExchangeRate, &t.Notes, &t.Date, &t.Status, &t.InstallmentPlanID,
		&t.RecurringPaymentID, &t.PaymentNumber,
	)
}
