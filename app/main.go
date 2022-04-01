package main

import (
	// "log"
	"net/http"
	"portfolio/handler"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	userkey = "user"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/login", handler.UserLogin())
	// r.GET("/ping", handler.TestApi())
	r.POST("/addskills", handler.AddSkills())
	r.GET("/getskills", handler.GETAllSkills())
	r.DELETE("/deleteskill/:id", handler.SkillDelete())
	r.POST("/addservice", handler.AddServices())
	r.GET("/getservices", handler.GETAllService())
	r.GET("/getexperience", handler.GetAllExperience())
	r.POST("/addexperience", handler.ExperienceAdd())
	r.GET("/getmsg", handler.GetAllMessage())
	r.POST("/addmsg", handler.AddMessage())
	r.POST("/addexpense", handler.ExpenseAdd())
	r.Run()

	// r := engine()
	// r.Use(gin.Logger())
	// if err := engine().Run(":8080"); err != nil {
	// 	log.Fatal("Unable to start:", err)
	// }
}
func engine() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.POST("/login", login)
	// r.GET("/ping", handler.TestApi())

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/status", status)
	}
	return r
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "mahidulmoon" || password != "1114012833mM#" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}
func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
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
