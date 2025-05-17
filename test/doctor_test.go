package test

import (
	
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang_patient_portal/config"
	"golang_patient_portal/models"
)

import "strconv"

func toStr(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}

func TestDoctorFetchPatient_Unauthorized(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest("GET", "/api/doctor/patients/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestDoctorUpdateNotes(t *testing.T) {
	// Step 1: Seed a test patient
	patient := models.Patient{
		Name:   "Note Patient",
		Age:    28,
		Gender: "male",
	}
	config.DB.Create(&patient)

	// Step 2: Prepare notes payload
	payload := map[string]string{
		"notes": "Patient is recovering well.",
	}
	body, _ := json.Marshal(payload)

	// Step 3: Create auth request using valid token
	req := authDoctorRequest("PUT", "/api/doctor/patients/"+toStr(patient.ID)+"/notes", body)
	w := httptest.NewRecorder()

	router := setupRouter()
	router.ServeHTTP(w, req)

	// Step 4: Assert
	assert.True(t, w.Code == 200 || w.Code == 204)
}