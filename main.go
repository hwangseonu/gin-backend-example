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

	post := r.Group("/posts")
	postRegister(post)

	r.Run()
}

func userRegister(group *gin.RouterGroup) {
	group.POST("", routes.SignUp)

	auth := group.Group("")
	{
		auth.Use(middlewares.AuthRequired("access"))
		auth.GET("", routes.UserInfo)
	}
}

func authRegister(group *gin.RouterGroup) {
	group.POST("", routes.Auth)

	refresh := group.Group("")
	{
		refresh.Use(middlewares.AuthRequired("refresh"))
		refresh.GET("/refresh", routes.Refresh)
	}
}

func postRegister(group *gin.RouterGroup) {
	group.GET("", routes.GetAllPost)
	group.GET("/:pid", routes.GetPost)
	auth := group.Group("")
	{
		auth.Use(middlewares.AuthRequired("access"))
		auth.POST("", routes.NewPost)
		auth.POST("/:pid/comments", routes.AddComment)
	}
}