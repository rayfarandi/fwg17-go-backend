package customer

import (
	"github.com/gin-gonic/gin"
	controllers_customer "github.com/rayfarandi/fwg17-go-backend/src/controllers/customer"
	"github.com/rayfarandi/fwg17-go-backend/src/middlewares"
)

func ProfileRouter(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())

	r.GET("", controllers_customer.GetProfile)
	r.PATCH("", controllers_customer.UpdateProfile)
}
