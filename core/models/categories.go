package models

import (
	"api/db"
	"database/sql"
	"fmt"
)

type SubCategory struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	CategoryID int    `json:"category_id"`
}

type Categories struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Type          string        `json:"type"`
	Icon          string        `json:"icon"`
	BudgetLimit   float64       `json:"budget_limit"`
	SubCategories []SubCategory `json:"subcategories,omitempty"`
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
        WITH CategoryData AS (
            SELECT c.*
            FROM Categories c
            ORDER BY c.id DESC
            LIMIT ? OFFSET ?
        )
        SELECT
            c.id,
            c.name,
            c.type,
            c.icon,
            c.budget_limit,
            sc.id as subcategory_id,
            sc.name as subcategory_name,
            sc.icon as subcategory_icon
        FROM CategoryData c
        LEFT JOIN SubCategories sc ON c.id = sc.category_id
        ORDER BY c.id DESC, sc.id ASC`)
}

func (c *Categories) Scan(rows *sql.Rows) error {
	var subCategory struct {
		ID   sql.NullInt64
		Name sql.NullString
		Icon sql.NullString
	}

	err := rows.Scan(
		&c.ID, &c.Name, &c.Type, &c.Icon, &c.BudgetLimit,
		&subCategory.ID, &subCategory.Name, &subCategory.Icon,
	)

	if err != nil {
		return err
	}

	// If we have a valid subcategory, add it to the slice
	if subCategory.ID.Valid {
		c.SubCategories = append(c.SubCategories, SubCategory{
			ID:         int(subCategory.ID.Int64),
			Name:       subCategory.Name.String,
			Icon:       subCategory.Icon.String,
			CategoryID: c.ID,
		})
	}

	return nil
}
