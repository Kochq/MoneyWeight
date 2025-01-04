package routes

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRecurringPayments(c *gin.Context) {
	GetEntities(c, func() *RecurringPayment {
		return &RecurringPayment{}
	})
}

func AddRecurringPayment(c *gin.Context) {
	AddEntity(c, func() *RecurringPayment {
		return &RecurringPayment{}
	})
}

func UpdateRecurringPayment(c *gin.Context) {
	UpdateEntity(c, func() *RecurringPayment {
		return &RecurringPayment{}
	})
}

func (RecurringPayment) TableName() string {
	return "RecurringPayments"
}

func (rp *RecurringPayment) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&rp.ID, &rp.Title, &rp.Amount, &rp.Frequency, &rp.StartDate, &rp.EndDate,
		&rp.IsActive, &rp.CategoryID, &rp.SubCategoryID,
	)
}

func (rp RecurringPayment) GetQuery() string {
	return fmt.Sprintf(`
    SELECT * 
    FROM %s
    ORDER BY ID DESC 
    LIMIT ? OFFSET ?`, rp.TableName())
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

	return rp.insert()
}

func (rp *RecurringPayment) insert() (int, error) {
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
    UPDATE RecurringPayment 
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
		&rp.StartDate, &rp.EndDate, &rp.Frequency,
	)
	if err != nil {
		return 0, err
	}

	idR, err := res.RowsAffected()

	return int(idR), err
}
