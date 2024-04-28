package routers

import (
	controllers_customer "github.com/rayfarandi/fwg17-go-backend/src/controllers/customer"
	"github.com/rayfarandi/fwg17-go-backend/src/middleware"

	"github.com/gin-gonic/gin"
)

func CombineCustomer(r *gin.RouterGroup) {
	AuthRouter(r.Group("/"))
	ProductsRouter(r.Group("/products"))

	authMiddleware, _ := middleware.Auth()
	r.Use(authMiddleware.MiddlewareFunc())

	ProfileRouter(r.Group("/profile"))
	HistoryOrderRouter(r.Group("/history-order"))
	r.GET("/order-products", controllers_customer.ListOrderProducts)
	r.POST("/checkout", controllers_customer.Checkout)
	r.GET("/data-size", controllers_customer.GetPriceSize)
	r.GET("/data-variant", controllers_customer.GetPriceVariant)
}
