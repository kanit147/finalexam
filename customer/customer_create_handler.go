package customer

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanit147/finalexam/database"
	"github.com/kanit147/finalexam/types"
)

func createCustomersHandler(c *gin.Context) {
	log.Println("createCustomersHandler:", time.Now())

	t := types.Customer{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row, errx := create(t)
	if errx != nil {
		c.JSON(http.StatusInternalServerError, errx)
		return
	}

	err := row.Scan(&t.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, t)
}

func create(t types.Customer) (*sql.Row, error) {
	return database.CreateCustomer(t)
}
