package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func OrderDetailsRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllOrderDetails)
	r.GET("/:id", controllers.DetailOrderDetails)
	r.POST("", controllers.CreateOrderDetails)
	r.PATCH("/:id", controllers.UpdateOrderDetails)
	r.DELETE("/:id", controllers.DeleteOrderDetails)
}
