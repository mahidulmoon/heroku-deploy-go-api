package main

import (
	"portfolio/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.TestApi())
	r.Run()
}
