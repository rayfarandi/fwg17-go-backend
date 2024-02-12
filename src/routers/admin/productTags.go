package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func ProductTagsRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllProductTags)
	r.GET("/:id", controllers_admin.DetailProductTags)
	r.POST("", controllers_admin.CreateProductTags)
	r.PATCH("/:id", controllers_admin.UpdateProductTags)
	r.DELETE("/:id", controllers_admin.DeleteProductTags)
}
