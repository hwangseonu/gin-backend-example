package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/routes"
	"log"
)

func main() {
	r := gin.Default()
	routes.InitRoutes(r)
	gin.SetMode(gin.DebugMode)
	log.Fatal(r.Run(":5000"))
}
