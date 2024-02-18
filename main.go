package main

import (
	// "log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/joho/godotenv"

	"github.com/rayfarandi/fwg17-go-backend/src/routers"
	"github.com/rayfarandi/fwg17-go-backend/src/services"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, Authorization"},
	}))
	//tes ENV
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("error load .env")
	// }
	// r.Static("/uploads", "./uploads")

	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &services.Response{
			Success: false,
			Message: "Resource not found restart on change yaa",
		})
	})
	r.Run("127.0.0.1:8888")
}
