package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func OrdersRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllOrders)
	r.GET("/:id", controllers.DetailOrder)
	r.POST("", controllers.CreateOrders)
	r.PATCH("/:id", controllers.UpdateOrders)
	r.DELETE("/:id", controllers.DeleteOrders)
}
