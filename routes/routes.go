package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// later: add routes for user registration, login, and booking management
	server.POST("/users")

}
