package handler

import (
	"net/http"
	"portfolio/models"

	"github.com/gin-gonic/gin"
)

func TestApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func AddSkills() gin.HandlerFunc {
	return func(c *gin.Context) {
		var skill models.Skills

		err := c.ShouldBindJSON(&skill)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error to bind": err,
			})
		} else {
			err := skill.Add()
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error to bind": err,
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{
					"success": "skills added successfully",
				})
			}
		}
	}
}
