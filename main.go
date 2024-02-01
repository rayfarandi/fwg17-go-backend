package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rayfarandi/fwg17-go-backend/src/routers"
)

func main() {
	r := gin.Default()
	routers.Combine(r)
	r.Run()
}
