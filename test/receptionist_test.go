package test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceptionistGetPatients_Unauthorized(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest("GET", "/api/receptionist/patients", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestReceptionistAddPatient(t *testing.T) {
	router := setupRouter()

	patient := map[string]interface{}{
		"name":   "Test User",
		"age":    30,
		"gender": "male",
	}
	body, _ := json.Marshal(patient)

	req := authRequest("POST", "/api/receptionist/patients", body)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.True(t, w.Code == 200 || w.Code == 201)
}

func TestReceptionistUpdatePatient(t *testing.T) {
	router := setupRouter()

	update := map[string]interface{}{
		"name":   "Updated Name",
		"age":    40,
		"gender": "female",
	}
	body, _ := json.Marshal(update)

	req := authRequest("PUT", "/api/receptionist/patients/1", body)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.True(t, w.Code == 200 || w.Code == 204)
}

func TestReceptionistDeletePatient(t *testing.T) {
	router := setupRouter()

	req := authRequest("DELETE", "/api/receptionist/patients/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.True(t, w.Code == 200 || w.Code == 204)
}
