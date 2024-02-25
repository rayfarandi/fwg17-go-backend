package routers

import (
	"github.com/rayfarandi/fwg17-go-backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func ForgotPasswordRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllforgotPassword)
	r.GET("/:id", controllers.DetailForgotPassword)
	r.POST("", controllers.CreateForgotPassword)
	r.PATCH("/:id", controllers.UpdateForgotPassword)
	r.DELETE("/:id", controllers.DeleteForgotPassword)
}
