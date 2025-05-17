package controllers

import (
	"net/http"
	"golang_patient_portal/config"
	"golang_patient_portal/models"
	"github.com/gin-gonic/gin"
	"fmt"
)

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Received patient: %+v\n", patient)
	config.DB.Create(&patient)
	c.JSON(http.StatusCreated, patient)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func UpdatePatient(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

func DeletePatient(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")

	if err := config.DB.Delete(&patient, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
