package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/routers/admin"
)

func Combine(r *gin.Engine) {
	AuthRouter((r.Group("/auth")))
	admin.Combine(r.Group("/admin"))

}
