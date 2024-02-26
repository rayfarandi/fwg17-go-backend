package main

import (
	"log"
	"net/http"

	"github.com/rayfarandi/fwg17-go-backend/src/routers"
	"github.com/rayfarandi/fwg17-go-backend/src/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, Authorization"},
	}))
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r.Static("/uploads", "./uploads")
	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Resource not found",
		})
	})
	// r.Run("127.0.0.1:8888")
	r.Run(":8888")
}
