package customer

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanit147/finalexam/database"
)

func deleteCustomersHandler(c *gin.Context) {
	log.Println("deleteCustomersHandler:", time.Now())
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	errx := deleteById(id)
	if errx != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	log.Println("delete customer id:", id)
	// t := &Message{Message: "customer deleted"}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}

func deleteById(id int) error {
	return database.DeleteById(id)
}
