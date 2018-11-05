package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend/routes"
)

func main() {
	r := gin.Default()

	user := r.Group("/users")
	userRegister(user)

	r.Run()
}

func userRegister(group *gin.RouterGroup) {
	group.POST("", routes.SignUp)
}