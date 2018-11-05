package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend/middlewares"
	"github.com/hwangseonu/gin-backend/routes"
)

func main() {
	r := gin.Default()

	user := r.Group("/users")
	userRegister(user)

	auth := r.Group("/auth")
	authRegister(auth)

	r.Run()
}

func userRegister(group *gin.RouterGroup) {
	group.POST("", routes.SignUp)
}

func authRegister(group *gin.RouterGroup) {
	group.POST("", routes.Auth)

	refresh := group.Group("")
	{
		refresh.Use(middlewares.AuthRequired("refresh"))
		refresh.GET("/refresh", routes.Refresh)
	}
}