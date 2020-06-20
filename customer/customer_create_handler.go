package customer

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanit147/finalexam/database"
)

func createCustomersHandler(c *gin.Context) {
	log.Println("createCustomersHandler:", time.Now())

	/////////////////////
	createTable := `CREATE TABLE IF NOT EXISTS CUSTOMERS (
		ID SERIAL PRIMARY KEY,
		NAME TEXT,
		EMAIL TEXT,
		STATUS TEXT
	);
		`
	database.Con().Exec(createTable)
	/////////////////////

	t := Customer{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := database.Con().QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id", t.Name, t.Email, t.Status)

	err := row.Scan(&t.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, t)
}
