package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func ProductRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllProduct)
	r.GET("/:id", controllers_admin.DetailProduct)
	r.POST("", controllers_admin.CreateProduct)
	r.PATCH("/:id", controllers_admin.UpdateProduct)
	r.DELETE("/:id", controllers_admin.DeleteProduct)
}
