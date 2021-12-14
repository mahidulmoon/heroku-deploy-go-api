package handler

import (
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
			c.JSON(400, gin.H{
				"error to bind": err,
			})
		} else {
			err := skill.Add()
			if err != nil {
				c.JSON(500, gin.H{
					"error to bind": err,
				})
			} else {
				c.JSON(201, gin.H{
					"success": "skills added successfully",
				})
			}
		}
	}
}

func GETAllSkills() gin.HandlerFunc {
	return func(c *gin.Context) {
		skillsData, err := models.GetAllSkills()
		if err != nil {
			c.JSON(200, gin.H{
				"results": "not found",
			})
		} else {
			c.JSON(500, gin.H{
				"results": skillsData,
			})
		}
	}
}
