package routes

import (
	"api/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Entity interface {
	TableName() string
	GetSelectQuery() string
	Scan(rows *sql.Rows) error

	Create() (int, error)
	Insert() (int, error)
	SetEntity(id int) (int, error)
	DeleteEntity(id int) (int, error)
}

// Generic function | T is an Entity | newT is a contructor for T
func GetEntities[T Entity](c *gin.Context, newT func() T) {
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	// This is just to get the table name
	entity := newT()

	rows, err := db.DB.Query(entity.GetSelectQuery(), limit, offset)
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

func AddEntity[T Entity](c *gin.Context, newT func() T) {
	var body = newT()
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request body",
		})
		return
	}

	id, err := body.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"id":     id,
	})
}

func UpdateEntity[T Entity](c *gin.Context, newT func() T) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Invalid ID",
		})
	}

	var body = newT()
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request body",
		})
		return
	}

	rowsAffected, err := body.SetEntity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// It's working perfectly fine...
	// it wont affect any rows if it can't change anything
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  body.TableName() + " not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf(body.TableName()+" with ID %d updated", id),
	})
}

func RemoveEntity[T Entity](c *gin.Context, newT func() T) {
	entity := newT()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Invalid ID",
		})
	}

	id, err = entity.DeleteEntity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf(entity.TableName()+" with ID %d deleted", id),
	})
}
