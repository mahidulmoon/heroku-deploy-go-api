package handler

import (
	"portfolio/models"
	"strconv"

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
			c.JSON(500, gin.H{
				"results": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"results": skillsData,
			})
		}
	}
}

func SkillDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := models.DeleteSkill(ID)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Skill deleted successfully.",
			})
		}
	}
}

func AddServices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var service models.Services

		err := c.ShouldBindJSON(&service)

		if err != nil {
			c.JSON(400, gin.H{
				"error to bind": err,
			})
		} else {
			err := service.Add()
			if err != nil {
				c.JSON(500, gin.H{
					"error to bind": err,
				})
			} else {
				c.JSON(201, gin.H{
					"success": "service added successfully",
				})
			}
		}
	}
}

func GETAllService() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceData, err := models.GetServices()
		if err != nil {
			c.JSON(500, gin.H{
				"results": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"results": serviceData,
			})
		}
	}
}

func GetAllExperience() gin.HandlerFunc {
	return func(c *gin.Context) {
		experience, err := models.GetAllExperience()
		if err != nil {
			c.JSON(500, gin.H{
				"results": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"results": experience,
			})
		}
	}
}

func ExperienceAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var exp models.Experience

		err := c.ShouldBindJSON(&exp)

		if err != nil {
			c.JSON(400, gin.H{
				"error to bind": err,
			})
		} else {
			err := exp.Add()
			if err != nil {
				c.JSON(500, gin.H{
					"error to bind": err,
				})
			} else {
				c.JSON(201, gin.H{
					"success": "service added successfully",
				})
			}
		}
	}
}

func GetAllMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		msg, err := models.GetAllMessage()
		if err != nil {
			c.JSON(500, gin.H{
				"results": "not found",
				"error":   err,
			})
		} else {
			c.JSON(200, gin.H{
				"results": msg,
			})
		}
	}
}
