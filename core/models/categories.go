package models

import (
	"api/db"
	"database/sql"
	"fmt"
)

type Categories struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Icon        string `json:"icon"`
	BudgetLimit float64 `json:"budget_limit"`
}

func (c *Categories) Create() (int, error) {
	id, err := c.Insert()
	if err != nil {
		return 0, fmt.Errorf("error creating category: %v", err)
	}

	return int(id), nil
}

func (c *Categories) Insert() (int, error) {
	query := `
    INSERT INTO Categories (name, type, icon, budget_limit) VALUES (?, ?, ?, ?)`

	result, err := db.DB.Exec(query, &c.Name, &c.Type, &c.Icon, &c.BudgetLimit)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (c Categories) SetEntity(id int) (int, error) {
	query := `
    UPDATE Categories
    SET name = ?,
    type = ?,
    icon = ?,
    budget_limit = ?
    WHERE id = ?`

	res, err := db.DB.Exec(query, &c.Name, &c.Type, &c.Icon, &c.BudgetLimit, id)
	if err != nil {
		return 0, err
	}

	idR, err := res.RowsAffected()

	return int(idR), err
}

func (c *Categories) DeleteEntity(id int) (int, error) {
	query := "DELETE FROM Categories WHERE id = ?"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("Error deleting category %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error trying to retrieve rows affected")
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("Couldnt remove category %d", id)
	}
	return id, nil
}

func (Categories) TableName() string {
	return "Categories"
}

func (i Categories) GetSelectQuery() string {
	return fmt.Sprintf(`
        SELECT *
        FROM %s
        ORDER BY ID DESC
        LIMIT ? OFFSET ?`, i.TableName())
}

func (c *Categories) Scan(rows *sql.Rows) error {
	return rows.Scan(&c.ID, &c.Name, &c.Icon, &c.Type, &c.BudgetLimit)
}
