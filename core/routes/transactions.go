package routes

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

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

func UpdateTransaction(c *gin.Context) {
	UpdateEntity(c, func() *Transaction {
		return &Transaction{}
	})
}

func RemoveTransaction(c *gin.Context) {
	RemoveEntity(c, func() *Transaction {
		return &Transaction{}
	})
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
		&t.ID, &t.Title, &t.Amount, &t.CategoryID, &t.SubCategoryID,
		&t.Currency, &t.PaymentMethod, &t.ExchangeRate, &t.Notes, &t.Date,
		&t.InstallmentPlanID, &t.RecurringPaymentID, &t.PaymentNumber,
	)
}
