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

func getCustomersHandler(c *gin.Context) {
	log.Println("getCustomersHandler:", time.Now())
	status := c.Query("status")

	rows, err := getAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	customers := []*types.Customer{}
	for rows.Next() {
		t := &types.Customer{}

		err := rows.Scan(&t.ID, &t.Name, &t.Email, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		customers = append(customers, t)
	}

	tt := []*types.Customer{}

	for _, item := range customers {
		if status != "" {
			if item.Status == status {
				tt = append(tt, item)
			}
		} else {
			tt = append(tt, item)
		}
	}

	c.JSON(http.StatusOK, tt)
}

func getAll() (*sql.Rows, error) {
	return database.GetCustomerAll()
}
