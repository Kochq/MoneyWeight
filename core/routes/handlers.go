package routes

import (
	"api/db"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Entity interface {
	TableName() string
	Scan(rows *sql.Rows) error
    GetQuery() string
}

// Generic function | T is an Entity | newT is a contructor for T
func GetEntities[T Entity](c *gin.Context, newT func() T) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	// This is just to get the table name
	entity := newT()

	rows, err := db.DB.Query(entity.GetQuery(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var entities []T
	for rows.Next() {
		entity := newT()
		if err := entity.Scan(rows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
			return
		}
		entities = append(entities, entity)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   entities,
	})
}