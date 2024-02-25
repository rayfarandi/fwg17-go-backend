package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRatingsRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllProductRatings)
	r.GET("/:id", controllers.DetailProductRatings)
	r.POST("", controllers.CreateProductRatings)
	r.PATCH("/:id", controllers.UpdatePrductRatings)
	r.DELETE("/:id", controllers.DeleteProductRatings)
}
