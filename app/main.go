package main

import (
	"portfolio/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	// r.GET("/ping", handler.TestApi())
	r.POST("/addskills", handler.AddSkills())
	r.GET("/getskills", handler.GETAllSkills())
	r.POST("/addservice", handler.AddServices())
	r.GET("/getservices", handler.GETAllService())
	r.GET("/getexperience", handler.GetAllExperience())
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, X-Auth-Secret, Uid, Aid, CToken")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
