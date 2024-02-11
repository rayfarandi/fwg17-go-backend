package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func ProductVariantRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllProductVariant)
	r.GET("/:id", controllers_admin.DetailProductVariant)
	r.POST("", controllers_admin.CreateProductVariant)
	r.PATCH("/:id", controllers_admin.UpdateProductVariant)
	r.DELETE("/:id", controllers_admin.DeleteProductVariant)
}
