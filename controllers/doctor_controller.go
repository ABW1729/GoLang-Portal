package controllers

import (
	"net/http"
	"golang_patient_portal/config"
	"golang_patient_portal/models"

	"github.com/gin-gonic/gin"
)

func GetPatientByID(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func UpdatePatientNotes(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var input struct {
		Notes string `json:"notes"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient.Notes = input.Notes
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}
