package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func PromoRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllPromo)
	r.GET("/:id", controllers.DetailPromo)
	r.POST("", controllers.CreatePromo)
	r.PATCH("/:id", controllers.UpdatePromo)
	r.DELETE("/:id", controllers.DeletePromo)
}
