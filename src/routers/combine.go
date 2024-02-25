package routers

import (
	admin "github.com/rayfarandi/fwg17-go-backend/src/routers/admin"
	customer "github.com/rayfarandi/fwg17-go-backend/src/routers/customer"

	"github.com/gin-gonic/gin"
)

func Combine(r *gin.Engine) {
	admin.CombineAdmin(r.Group("/admin"))
	customer.CombineCustomer(r.Group("/"))
}
