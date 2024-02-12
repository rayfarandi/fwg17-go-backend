package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func ProductRatingsRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllProductRatings)
	r.GET("/:id", controllers_admin.DetailProductRatings)
	r.POST("", controllers_admin.CreateProductRatings)
	r.PATCH("/:id", controllers_admin.UpdateProductRatings)
	r.DELETE("/:id", controllers_admin.DeleteProductRatings)
}
