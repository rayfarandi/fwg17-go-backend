package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"
	controllers_customer "github.com/rayfarandi/fwg17-go-backend/src/controllers/customer"

	"github.com/gin-gonic/gin"
)

func HistoryOrderRouter(r *gin.RouterGroup) {
	r.GET("", controllers_customer.ListAllOrders)
	r.GET("/:id", controllers.DetailOrder)
}
