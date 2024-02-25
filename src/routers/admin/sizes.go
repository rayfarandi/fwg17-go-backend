package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func SizesRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllSizes)
	r.GET("/:id", controllers.DetailSizes)
	r.POST("", controllers.CreateSize)
	r.PATCH("/:id", controllers.UpdateSizes)
	r.DELETE("/:id", controllers.DeleteSizes)
}
