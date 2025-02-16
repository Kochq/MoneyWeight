package models

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"
)

type Account struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	CurrentBalance float64 `json:"current_balance"`
	Currency       string  `json:"currency"`
	Institution    string  `json:"institution"`
	IsActive       bool    `json:"is_active"`
	CreatedAt      string  `json:"created_at"` // Podríamos usar time.Time
}

func (a *Account) Create() (int, error) {
	if a.CreatedAt == "" {
		a.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	}
	return a.Insert()
}

func (a *Account) Insert() (int, error) {
	query := `
    INSERT INTO Accounts (
    name, type, current_balance, currency, institution, is_active, created_at
    ) VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query,
		a.Name, a.Type, a.CurrentBalance, a.Currency, a.Institution,
		a.IsActive, a.CreatedAt,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (a Account) SetEntity(id int) (int, error) {
	query := `
    UPDATE Accounts
    SET name = ?,
    type = ?,
    current_balance = ?,
    currency = ?,
    institution = ?
    is_active = ?
    created_at = ?
    WHERE id = ?
    `

	res, err := db.DB.Exec(query,
		a.Name, a.Type, a.CurrentBalance, a.Currency, a.Institution,
		a.IsActive, a.CreatedAt, id,
	)
	if err != nil {
		return 0, err
	}

	idR, err := res.RowsAffected()

	return int(idR), err
}

func (t *Account) DeleteEntity(id int) (int, error) {
	query := "DELETE FROM Accounts WHERE id = ?"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("Error deleting account %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error trying to retrieve rows affected")
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("Couldnt remove account %d", id)
	}

	return id, nil
}

func (Account) TableName() string {
	return "Accounts"
}

func (t Account) GetSelectQuery() string {
	return fmt.Sprintf(`
        SELECT *
        FROM %s
        LIMIT ? OFFSET ?`, t.TableName())
}

func (a *Account) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&a.ID, &a.Name, &a.Type, &a.CurrentBalance, &a.Currency, &a.Institution,
		&a.IsActive, &a.CreatedAt,
	)
}
