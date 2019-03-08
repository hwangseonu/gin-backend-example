package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/controllers"
	"github.com/hwangseonu/gin-backend-example/server/middlewares"
	"github.com/hwangseonu/gin-backend-example/server/requests"
)

func InitAuthRoute(e *gin.RouterGroup) {
	signIn := e.Use(middlewares.JsonRequired(&requests.SignInRequest{}))
	signIn.POST("", controllers.SignIn)
}
