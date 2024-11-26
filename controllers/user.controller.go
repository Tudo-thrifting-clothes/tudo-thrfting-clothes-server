package controllers

import (
	"net/http"

	"tudo-thrifting-server/services"

	"github.com/gin-gonic/gin"
)

// GetUsers handles GET requests to /users
func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID handles GET requests to /users/:id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
