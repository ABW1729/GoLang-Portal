package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang_patient_portal/models"
	"golang_patient_portal/utils"
)

func Login(c *gin.Context) {
	var credentials models.LoginRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := models.AuthenticateUser(credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	token, _ := utils.GenerateJWT(user)
	c.JSON(http.StatusOK, gin.H{"token": token, "role": user.Role})
}
