package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func ProductCategoriesRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllProductCategories)
	r.GET("/:id", controllers_admin.DetailProductCategories)
	r.POST("", controllers_admin.CreateProductCategories)
	r.PATCH("/:id", controllers_admin.UpdateProductCategories)
	r.DELETE("/:id", controllers_admin.DeleteProductCategories)
}
