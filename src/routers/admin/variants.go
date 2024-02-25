package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func VariantsRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllVariants)
	r.GET("/:id", controllers.DetailVariants)
	r.POST("", controllers.CreateVariants)
	r.PATCH("/:id", controllers.UpdateVariants)
	r.DELETE("/:id", controllers.DeleteVariants)
}
