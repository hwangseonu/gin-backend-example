package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/controllers"
	"github.com/hwangseonu/gin-backend-example/server/middlewares"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/security"
)

func InitUserRoute(e *gin.RouterGroup) {
	signUp := e.Group("")
	signUp.Use(middlewares.JsonRequired(&requests.SignUpRequest{}))
	signUp.POST("", controllers.SignUp)
	getUser := e.Group("")
	getUser.Use(middlewares.AuthRequired(security.ACCESS, "ROLE_USER"))
	getUser.GET("", controllers.GetUser)
}
