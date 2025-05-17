package test
import _ "github.com/joho/godotenv/autoload" 
import (
	"os"
	"testing"

	"golang_patient_portal/config"
	"golang_patient_portal/models"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_NAME", "hospital_test")
	os.Setenv("DB_USER", "aniket")
	os.Setenv("DB_PASSWORD", "1234")
	os.Setenv("DB_PORT","5432")
	os.Setenv("DB_HOST","localhost")
	config.ConnectDatabase()

	// Auto-migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Patient{})

	// Clean and seed
	config.DB.Exec("TRUNCATE patients, users RESTART IDENTITY CASCADE")

	// Create sample receptionist
	config.DB.Create(&models.User{
		Username: "receptionist1",
		Password: "password123", 
		Role:     "receptionist",
	})

	config.DB.Create(&models.User{
    Username: "doctor1",
    Password: "password123", 
    Role:     "doctor",
})

	code := m.Run()
	os.Exit(code)
}
