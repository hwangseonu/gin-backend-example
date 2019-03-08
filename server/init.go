package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/routes"
)

func GenerateApp() *gin.Engine {
	r := gin.Default()
	routes.InitRoutes(r)
	gin.SetMode(gin.DebugMode)
	return r
}

