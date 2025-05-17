package services

import (
	"golang_patient_portal/models"
	"golang_patient_portal/repository"
)

func AddNewPatient(p *models.Patient) error {
	return repository.CreatePatient(p)
}

func FetchAllPatients() ([]models.Patient, error) {
	return repository.GetAllPatients()
}

func ModifyPatient(id uint, data *models.Patient) error {
	return repository.UpdatePatient(id, data)
}

func RemovePatient(id uint) error {
	return repository.DeletePatient(id)
}
