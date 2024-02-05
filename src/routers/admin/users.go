package admin

import (
	"github.com/gin-gonic/gin"
	controllers_admin "github.com/rayfarandi/fwg17-go-backend/src/controllers/admin"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("", controllers_admin.ListAllUsers)
	r.GET("/:id", controllers_admin.DetailUser)
	r.POST("", controllers_admin.CreateUser)
	r.PATCH("/:id", controllers_admin.UpdateUser)
	r.DELETE("/:id", controllers_admin.DeleteUser)
}
