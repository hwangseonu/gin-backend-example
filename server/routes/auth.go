package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/controllers"
	"github.com/hwangseonu/gin-backend-example/server/middlewares"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/security"
)

func InitAuthRoute(e *gin.RouterGroup) {
	signIn := e.Group("")
	signIn.Use(middlewares.JsonRequired(&requests.SignInRequest{}))
	signIn.POST("", controllers.SignIn)

	refresh := e.Group("/refresh")
	refresh.Use(middlewares.AuthRequired(security.REFRESH))
	refresh.GET("", controllers.Refresh)
}
