package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func PromoRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllPromo)
	r.GET("/:id", controllers_admin.DetailPromo)
	r.POST("", controllers_admin.CreatePromo)
	r.PATCH("/:id", controllers_admin.UpdatePromo)
	r.DELETE("/:id", controllers_admin.DeletePromo)
}
