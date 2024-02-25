package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProductsRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllProducts)
	r.GET("/:id", controllers.DetailProducts)
}
