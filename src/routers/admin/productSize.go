package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func ProductSizeRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllProductSize)
	r.GET("/:id", controllers_admin.DetailProductSize)
	r.POST("", controllers_admin.CreateProductSize)
	r.PATCH("/:id", controllers_admin.UpdateProductSize)
	r.DELETE("/:id", controllers_admin.DeleteProductSize)
}
