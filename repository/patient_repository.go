package repository

import (
	"golang_patient_portal/config"
	"golang_patient_portal/models"
)

func CreatePatient(patient *models.Patient) error {
	return config.DB.Create(patient).Error
}

func GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	err := config.DB.Find(&patients).Error
	return patients, err
}

func UpdatePatient(id uint, updatedData *models.Patient) error {
	return config.DB.Model(&models.Patient{}).Where("id = ?", id).Updates(updatedData).Error
}

func DeletePatient(id uint) error {
	return config.DB.Delete(&models.Patient{}, id).Error
}
