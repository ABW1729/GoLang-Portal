Golang Patient Portal 

A backend web application built using Golang (Gin framework) with PostgreSQL. It provides role-based portals for Receptionists and Doctors to manage and access patient information.

FEATURES
- JWT-based Authentication
- Single Login API for Receptionist and Doctor
- Receptionist: Create, View, Update, and Delete Patients
- Doctor: View individual patients and update medical notes
- Role-based Authorization Middleware
- Modular architecture (Controller, Service, Repository pattern)
- Swagger/Postman API Documentation
- Unit Tests

PROJECT STRUCTURE
project/
├── cmd/                    # Application entrypoint
├── config/                 # DB config
├── controllers/            # Route handlers
├── middlewares/            # Auth middleware
├── models/                 # GORM models
├── repository/             # Data access layer
├── routes/                 # Route definitions
├── services/               # Business logic
├── utils/                  # Helper functions (e.g., JWT)
├── test/                   # Unit tests
├── docs/                   # Postman collection
├── go.mod / go.sum         # Go module files
└── README.md

SETUP INSTRUCTIONS
1. Clone the Repository
   git clone https://github.com/your-username/golang-patient-portal.git
   cd golang-patient-portal

2. Configure PostgreSQL
   sudo -u postgres psql
   CREATE DATABASE hospital;
   \q

3. Set the Environment Variable
   export DATABASE_URL="host=localhost user=postgres password=yourpassword dbname=hospital port=5432 sslmode=disable"

4. Install Dependencies
   go mod tidy

5. Run the App
   go run cmd/main.go

RUN UNIT TESTS
   go test ./test/...

API ENDPOINTS

Auth
- POST /api/login

Receptionist
- POST /api/receptionist/patients
- GET /api/receptionist/patients
- PUT /api/receptionist/patients/:id
- DELETE /api/receptionist/patients/:id

Doctor
- GET /api/doctor/patients/:id
- PUT /api/doctor/patients/:id/notes

API DOCUMENTATION
- Import docs/postman_collection.json into Postman

