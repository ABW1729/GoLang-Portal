package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"golang_patient_portal/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)
	return r
}

func getValidToken() string {
	router := setupRouter()

	creds := map[string]string{
		"username": "receptionist1",
		"password": "password123",
	}
	body, _ := json.Marshal(creds)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var res map[string]string
	json.Unmarshal(w.Body.Bytes(), &res)
	return res["token"]
}

func authRequest(method, url string, body []byte) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+getValidToken())
	req.Header.Set("Content-Type", "application/json")
	return req
}

func getDoctorToken() string {
	router := setupRouter()

	loginPayload := map[string]string{
		"username": "doctor1",
		"password": "password123",
	}
	body, _ := json.Marshal(loginPayload)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var res map[string]string
	json.Unmarshal(w.Body.Bytes(), &res)
	return res["token"]
}

func authDoctorRequest(method, url string, body []byte) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+getDoctorToken())
	req.Header.Set("Content-Type", "application/json")
	return req
}

