package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/controllers"
	"github.com/hwangseonu/gin-backend-example/server/middlewares"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/security"
)

func InitPostRoute(e *gin.RouterGroup) {
	create := e.Group("")
	create.Use(middlewares.AuthRequired(security.ACCESS, "ROLE_USER"))
	create.Use(middlewares.JsonRequired(&requests.CreatePostRequest{}))
	create.POST("", controllers.CreatePost)

	get := e.Group("/:post_id")
	get.GET("", controllers.GetPost)

	update := e.Group("/:post_id")
	update.Use(middlewares.AuthRequired(security.ACCESS, "ROLE_USER"))
	update.Use(middlewares.JsonRequired(&requests.CreatePostRequest{}))
	update.PATCH("", controllers.UpdatePost)

	del := e.Group("/:post_id")
	del.Use(middlewares.AuthRequired(security.ACCESS, "ROLE_USER"))
	del.DELETE("", controllers.DeletePost)
}
