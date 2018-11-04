package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend/server/controllers"
)

func main() {
	r := gin.Default()
	user := r.Group("/users")
	{
		user.POST("", controllers.CreateUser)
	}
	r.Run()
}