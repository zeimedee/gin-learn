package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zeimedee/gin-learn/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", handlers.Hello)
	r.GET("/test", handlers.Test)
	r.POST("/test", handlers.GetBody)
	r.GET("/users", handlers.GetUsers)
	r.GET("/user", handlers.User)
	return r
}
func main() {
	r := SetupRouter()
	r.Run(":3000")
}
