package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/controllers"
	"github.com/hwangseonu/gin-backend-example/server/middlewares"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/security"
)

func InitCommentRoutes(e *gin.RouterGroup) {
	add := e.Group("")
	add.Use(middlewares.JsonRequired(&requests.AddCommentRequest{}))
	add.Use(middlewares.AuthRequired(security.ACCESS, "ROLE_USER"))
	add.POST("", controllers.AddComment)
}
