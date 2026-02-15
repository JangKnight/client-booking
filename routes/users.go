package routes

import (
	"net/http"

	"example.com/client-booking/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user_id": user.ID})
}
