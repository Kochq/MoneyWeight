package models

import (
	"api/db"
	"database/sql"
	"fmt"
)

type SubCategories struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Icon        string  `json:"icon"`
	BudgetLimit float64 `json:"budget_limit"`
}

func (c *SubCategories) Create() (int, error) {
	id, err := c.Insert()
	if err != nil {
		return 0, fmt.Errorf("error creating SubCategory: %v", err)
	}

	return int(id), nil
}

func (c *SubCategories) Insert() (int, error) {
	query := `
    INSERT INTO SubCategories (name, category_id, icon, budget_limit) VALUES (?, ?, ?, ?)`

	result, err := db.DB.Exec(query, &c.Name, &c.CategoryID, &c.Icon, &c.BudgetLimit)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (c SubCategories) SetEntity(id int) (int, error) {
	query := `
    UPDATE SubCategories
    SET name = ?,
    category_id = ?,
    icon = ?,
    budget_limit = ?
    WHERE id = ?`

	res, err := db.DB.Exec(query, &c.CategoryID, &c.Name, &c.Icon, &c.BudgetLimit, id)
	if err != nil {
		return 0, err
	}

	idR, err := res.RowsAffected()

	return int(idR), err
}

func (c *SubCategories) DeleteEntity(id int) (int, error) {
	query := "DELETE FROM SubCategories WHERE id = ?"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("Error deleting category %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Error trying to retrieve rows affected")
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("Couldnt remove SubCategory %d", id)
	}
	return id, nil
}

func (SubCategories) TableName() string {
	return "SubCategories"
}

func (i SubCategories) GetSelectQuery() string {
	return fmt.Sprintf(`
        SELECT *
        FROM %s
        ORDER BY ID DESC
        LIMIT ? OFFSET ?`, i.TableName())
}

func (c *SubCategories) Scan(rows *sql.Rows) error {
	return rows.Scan(&c.ID, &c.CategoryID, &c.Name, &c.Icon, &c.BudgetLimit)
}
