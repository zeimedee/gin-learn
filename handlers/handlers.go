package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeimedee/gin-learn/models"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world!",
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "endpoint works",
	})
}

// gin function that returns request body

func GetBody(c *gin.Context) {
	user := new(models.User)
	if err := c.BindJSON(user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
		"age":  user.Age,
	})
}

//gin functions that returns a list of users

func GetUsers(c *gin.Context) {
	users := []models.User{
		{
			Name: "John",
			Age:  30,
		},
		{
			Name: "Jane",
			Age:  25,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

//gin function that returns a single user

func User(c *gin.Context) {
	user := models.User{
		Name: "John Smith",
		Age:  30,
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}
