package handler

import (
	"portfolio/models"
	"portfolio/services"
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

func JWTValidation(tokenString string) bool {
	token, err := services.JWTAuthService().ValidateToken(tokenString)
	if err != nil {
		return false
	} else {
		if token.Valid {
			return true
		} else {
			return false
		}
	}
}

func AddSkills() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("Authorization")
		// result := JWTValidation(token)
		// if result != true {
		// 	c.JSON(400, gin.H{
		// 		"message": "JWT token is not valid",
		// 	})
		// } else {
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
		// }
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
		// token := c.GetHeader("Authorization")
		// result := JWTValidation(token)
		// if result != true {
		// 	c.JSON(400, gin.H{
		// 		"message": "JWT token is not valid",
		// 	})
		// } else {
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
		// }
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
		// token := c.GetHeader("Authorization")
		// result := JWTValidation(token)
		// if result != true {
		// 	c.JSON(400, gin.H{
		// 		"message": "JWT token is not valid",
		// 	})
		// } else {
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
		// }
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

func AddMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var modelData models.MessageMe

		err := c.ShouldBindJSON(&modelData)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "can not bind data into json format.",
			})
		} else {
			err := modelData.AddMsg()
			if err != nil {
				c.JSON(400, gin.H{
					"message": "can not insert message please try again.",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "Your message has been stored in Database.",
				})
			}
		}
	}
}

func ExpenseAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var expense models.Expense

		err := c.ShouldBindJSON(&expense)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "can not bind data into json format.",
			})
		} else {
			err := expense.AddExpense()
			if err != nil {
				c.JSON(400, gin.H{
					"message": "can not insert expense please try again.",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "Your expense has been stored in Database.",
				})
			}
		}
	}
}

func AllExpense() gin.HandlerFunc {
	return func(c *gin.Context) {
		exp, err := models.GetAllExpense()
		if err != nil {
			c.JSON(500, gin.H{
				"results": "data not found",
				"error":   err,
			})
		} else {
			c.JSON(200, gin.H{
				"results": exp,
			})
		}
	}
}
