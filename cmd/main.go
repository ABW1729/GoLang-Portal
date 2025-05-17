package main
import _ "github.com/joho/godotenv/autoload" 

import (
    "golang_patient_portal/config"
    "golang_patient_portal/routes"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    config.ConnectDatabase()

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    routes.RegisterRoutes(r)

    r.Run(":8080")
}
