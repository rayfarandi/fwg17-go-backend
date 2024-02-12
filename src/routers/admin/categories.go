package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func CategoriesRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllCategories)
	r.GET("/:id", controllers_admin.DetailCategories)
	r.POST("", controllers_admin.CreateCategories)
	r.PATCH("/:id", controllers_admin.UpdateCategories)
	r.DELETE("/:id", controllers_admin.DeleteCategories)
}
