package models

import (
	"errors"
	"golang_patient_portal/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // doctor or receptionist
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthenticateUser(req LoginRequest) (*User, error) {
	var user User
	if err := config.DB.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}

type Patient struct {
	gorm.Model
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Notes  string `json:"notes"`
}
