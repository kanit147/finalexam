package customer

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanit147/finalexam/database"
	"github.com/kanit147/finalexam/types"
)

func getCustomerByIdHandler(c *gin.Context) {
	log.Println("getCustomerByIdHandler:", time.Now())
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	row, err := getById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	t := &types.Customer{}

	err = row.Scan(&t.ID, &t.Name, &t.Email, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, t)
}

func getById(id int) (*sql.Row, error) {
	return database.GetCustomerById(id)
}
