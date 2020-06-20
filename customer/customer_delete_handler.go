package customer

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanit147/finalexam/database"
)

func deleteCustomersHandler(c *gin.Context) {
	log.Println("deleteCustomersHandler:", time.Now())
	id := c.Param("id")
	stmt, err := database.Con().Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		log.Println("can't prepare delete statement", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		log.Println("can't execute delete statment", err)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	log.Println("delete customer id:", id)
	// t := &Message{Message: "customer deleted"}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
