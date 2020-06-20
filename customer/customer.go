package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/kanit147/finalexam/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/")

	api.Use(middleware.AuthorizationMdw)

	api.GET("customers", getCustomersHandler)
	api.GET("customers/:id", getCustomerByIdHandler)
	api.POST("customers", createCustomersHandler)
	api.PUT("customers/:id", updateCustomersHandler)
	api.DELETE("customers/:id", deleteCustomersHandler)

	return r
}
