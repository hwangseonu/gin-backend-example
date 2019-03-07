package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func (c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	gin.SetMode(gin.DebugMode)
	log.Fatal(r.Run(":5000"))
}
