Golang Patient Portal

A backend web application built using Golang (Gin framework) with PostgreSQL. It provides role-based portals for Receptionists and Doctors to manage and access patient information.

Features  
JWT-based Authentication  
Single Login API for Receptionist and Doctor  
Receptionist: Create, View, Update, and Delete Patients  
Doctor: View individual patients and update medical notes  
Role-based Authorization MiddlewareModular architecture (Controller, Service, Repository pattern)  
Swagger/Postman API Documentation
Unit Tests


Project Structure

project/  
├── cmd/                           # Application entrypoint  
├── config/                        # DB config  
├── controllers/                 # Route handlers  
├── middlewares/                 # Auth middleware  
├── models/                        # GORM models  
├── repository/                  # Data access layer    
├── routes/                        # Route definitions  
├── services/                      # Business logic  
├── utils/                         # Helper functions (e.g., JWT)  
├── test/                          # Unit tests  
├── docs/                          # Postman collection  
├── go.mod / go.sum              # Go module files  
└── README.md  


Setup Instructions
Clone the Repository
```bash
git clone https://github.com/your-username/golang-patient-portal.git
cd golang-patient-portal
```
Configure PostgreSQL
```bash
sudo -u postgres psql
CREATE DATABASE hospital;
\q
```

Set the Environment Variable
```bash
export DATABASE_URL="host=localhost user=postgres password=yourpassword dbname=hospital port=5432 sslmode=disable"
```
Install Dependencies
```bash
go mod tidy
```
Run the App
```bash
go run cmd/main.go
```
Run Unit Tests
```bash
go test ./test/...
```

API Endpoints  
Auth  
POST /api/loginReceptionist  
POST /api/receptionist/patients  
GET /api/receptionist/patients  
PUT /api/receptionist/patients/:id  
DELETE /api/receptionist/patients/:id    

Doctor  
GET /api/doctor/patients/:id  
PUT /api/doctor/patients/:id/notes  

API Documentation  
Import docs/postman_collection.json into Postman  
