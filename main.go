package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/routers"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func main() {
	r := gin.Default()
	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &services.Response{
			Success: false,
			Message: "Resource not found restart on change yaa",
		})
	})
	r.Run("127.0.0.1:8888")
}
