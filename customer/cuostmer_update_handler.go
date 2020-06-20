package customer

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanit147/finalexam/database"
	"github.com/kanit147/finalexam/types"
)

func updateCustomersHandler(c *gin.Context) {
	log.Println("updateCustomersHandler:", time.Now())
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	t := &types.Customer{ID: id}
	// map customer
	if err := c.ShouldBindJSON(t); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// update customer
	err = update(*t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, t)
}

func update(t types.Customer) error {
	return database.UpdateCustomer(t)
}
