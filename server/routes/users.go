package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/controllers"
	"github.com/hwangseonu/gin-backend-example/server/middlewares"
	"github.com/hwangseonu/gin-backend-example/server/requests"
)

func InitUserRoute(e *gin.RouterGroup) {
	signUp := e.Use(middlewares.JsonRequired(&requests.SignUpRequest{}))
	signUp.POST("", controllers.SignUp)
}
