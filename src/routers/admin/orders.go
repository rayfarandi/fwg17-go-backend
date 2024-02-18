package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func OrdersRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllOrders)
	r.GET("/:id", controllers_admin.DetailOrder)
	r.POST("", controllers_admin.CreateOrder)
	r.PATCH("/:id", controllers_admin.UpdateOrder)
	r.DELETE("/:id", controllers_admin.DeleteOrder)
}
