package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/rayfarandi/fwg17-go-backend/src/controllers/auth"
	"github.com/rayfarandi/fwg17-go-backend/src/middlewares"
)

func AuthRouter(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", controllers.Register)
	r.POST("/forgot-password", controllers.ForgotPassword)
}
