package routes

import (
	"golang_patient_portal/controllers"
	"golang_patient_portal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/login", controllers.Login)

	reception := api.Group("/receptionist")
	reception.Use(middlewares.AuthMiddleware("receptionist"))
	{
		reception.POST("/patients", controllers.CreatePatient)
		reception.GET("/patients", controllers.GetPatients)
		reception.PUT("/patients/:id", controllers.UpdatePatient)
		reception.DELETE("/patients/:id", controllers.DeletePatient)
	}

	doctor := api.Group("/doctor")
	doctor.Use(middlewares.AuthMiddleware("doctor"))
	{
		doctor.GET("/patients", controllers.GetPatients)
		doctor.GET("/patients/:id", controllers.GetPatientByID)
		doctor.PUT("/patients/:id/notes", controllers.UpdatePatientNotes)
	}
}
