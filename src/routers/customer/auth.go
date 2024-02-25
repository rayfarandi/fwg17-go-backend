package routers

import (
	controllers_customer "github.com/rayfarandi/fwg17-go-backend/src/controllers/customer"
	"github.com/rayfarandi/fwg17-go-backend/src/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	authMiddleware, _ := middleware.Auth()

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", controllers_customer.Register)
	r.POST("/forgot-password", controllers_customer.ForgotPassword)
}
