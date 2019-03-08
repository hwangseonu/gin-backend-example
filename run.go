package main

import (
	"github.com/hwangseonu/gin-backend-example/server"
	"log"
)
func main() {
	log.Fatal(server.GenerateApp().Run(":5000"))
}
