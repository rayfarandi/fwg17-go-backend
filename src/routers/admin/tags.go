package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func TagsRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllTags)
	r.GET("/:id", controllers.DetailTags)
	r.POST("", controllers.CreateTags)
	r.PATCH("/:id", controllers.UpdateTags)
	r.DELETE("/:id", controllers.DeleteTags)
}
