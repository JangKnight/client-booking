package main

import (
	"example.com/client-booking/db"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDB()
	server.Run(":8080")
}
