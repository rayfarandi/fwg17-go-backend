package routers

import (
	controllers_customer "github.com/rayfarandi/fwg17-go-backend/src/controllers/customer"
	"github.com/rayfarandi/fwg17-go-backend/src/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(r *gin.RouterGroup) {
	authMiddleware, _ := middleware.Auth()
	r.Use(authMiddleware.MiddlewareFunc())

	r.GET("", controllers_customer.GetProfile)
	r.PATCH("", controllers_customer.UpdateProfile)
}
