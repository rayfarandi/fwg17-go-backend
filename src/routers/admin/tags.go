package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func TagsRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllTags)
	r.GET("/:id", controllers_admin.DetailTags)
	r.POST("", controllers_admin.CreateTags)
	r.PATCH("/:id", controllers_admin.UpdateTags)
	r.DELETE("/:id", controllers_admin.DeleteTags)
}
