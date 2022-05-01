package handler

import (
	"fmt"
	"portfolio/models"
	"portfolio/services"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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

func JWTDecode(tokenString string) int64 {
	token, _ := jwt.Parse(tokenString, nil)
	if token == nil {
		// return nil, err
		fmt.Println("Invalid token", nil)
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	// fmt.Println("name: ", claims["user_id"])
	// claims are actually a map[string]interface{}
	userId, _ := strconv.ParseInt(fmt.Sprintf("%v", claims["user_id"]), 0, 64)
	return userId
}

func AddSkills() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		result := JWTValidation(token)
		if result != true {
			c.JSON(400, gin.H{
				"message": "JWT token is not valid",
			})
		} else {
			var skill models.Skills
			user_id := JWTDecode(token)

			err := c.ShouldBindJSON(&skill)

			if err != nil {
				c.JSON(400, gin.H{
					"error to bind": err,
				})
			} else {
				err := skill.Add(user_id)
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

func ExpenseFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter models.FilterExpense

		err := c.ShouldBindJSON(&filter)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "can not bind data into json format.",
			})
		} else {
			data, err := models.GetFilterExpense(filter.Year, filter.Month, filter.Day)
			if err != nil {
				c.JSON(400, gin.H{
					"message": "no data found",
					"error":   err,
				})
			} else {
				var total float64
				var total_in float64
				for value := range data {
					// fmt.Println("data: ", data[value].Price)
					if data[value].Type == "expense" {
						amount, err := strconv.ParseFloat(data[value].Price, 64)
						if err != nil {
							c.JSON(400, gin.H{
								"message": "float64 convertion error",
							})
						}
						total = total + amount
					} else {
						amount_in, err := strconv.ParseFloat(data[value].Price, 64)
						if err != nil {
							c.JSON(400, gin.H{
								"message": "float64 convertion error",
							})
						}
						total_in = total_in + amount_in
					}
				}
				c.JSON(200, gin.H{
					"message":         "data fetch successful",
					"data":            data,
					"total_amount":    total,
					"total_amount_in": total_in,
				})
			}

		}

	}
}

func GenerateMonthlyReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postvalues models.GenerateExpense
		err := c.ShouldBindJSON(&postvalues)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "can not bind data into json format.",
			})
		} else {
			msg, _ := postvalues.GetGenerateInfo()
			if msg == "no data found" || msg == "already data found" {
				c.JSON(400, gin.H{
					"message": msg,
				})
			} else {
				data, err := models.GetFilterExpense(postvalues.Year, postvalues.Month, postvalues.Monthdate)
				if err != nil {
					c.JSON(400, gin.H{
						"message": "no data found",
						"error":   err,
					})
				} else {
					var total float64
					var total_income float64

					for value := range data {
						if data[value].Type == "expense" {
							amount, err := strconv.ParseFloat(data[value].Price, 64)
							if err != nil {
								c.JSON(400, gin.H{
									"message": "float64 convertion error",
								})
							}
							total = total + amount
						} else {
							amount, err := strconv.ParseFloat(data[value].Price, 64)
							if err != nil {
								c.JSON(400, gin.H{
									"message": "float64 convertion error",
								})
							}
							total_income = total_income + amount
						}
					}
					if total == 0 {
						c.JSON(400, gin.H{
							"message": "No Data Found",
						})
					} else {
						err := postvalues.AddGenExp(fmt.Sprintf("%v", total), fmt.Sprintf("%v", total_income))
						if err != nil {
							c.JSON(400, gin.H{
								"message": "data generated unsuccessfull",
								"error":   err,
							})
						} else {
							err := models.DeleteGenExp(postvalues.Year, postvalues.Month)
							if err != nil {
								c.JSON(400, gin.H{
									"message": "data generated successfully but can not deleted from db",
								})
							} else {
								c.JSON(200, gin.H{
									"message": "data generated successfully",
								})
							}
						}
					}
				}
			}
		}
	}
}

func GetAllGenExp() gin.HandlerFunc {
	return func(c *gin.Context) {
		genexp, err := models.AllGeneData()
		if err != nil {
			c.JSON(500, gin.H{
				"results": "data not found",
				"error":   err,
			})
		} else {
			c.JSON(200, gin.H{
				"results": genexp,
			})
		}
	}
}
