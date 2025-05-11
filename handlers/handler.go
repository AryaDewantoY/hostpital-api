package handlers

import (
	"encoding/json"
	"hospital-api/config"
	"hospital-api/models"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var Secret = []byte("mysecret")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login endpoint
func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	_ = json.NewDecoder(r.Body).Decode(&creds)

	var user models.User
	// Akses DB dari package config
	err := config.DB.Where("email = ?", creds.Email).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Memeriksa password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) // Cukup gunakan err yang sudah ada
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString(config.Secret) // Menggunakan err yang sama
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	// Kirim token dalam respon
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// Middleware untuk auth
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// format: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return Secret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ==================== PATIENT ====================

// GetPatients godoc
// @Summary Get list of patients
// @Description Get all patients data
// @Tags patients
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Patient
// @Failure 401 {object} map[string]string
// @Router /api/patients [get]
func GetPatients(w http.ResponseWriter, r *http.Request) {
	patients := []models.Patient{
		{ID: 1, Name: "Mrs. Sheila Brakus", Age: 87, Gender: "Female", Address: ""},
		{ID: 2, Name: "Princess Dana Legros", Age: 65, Gender: "Male", Address: ""},
		{ID: 3, Name: "Dr. Janae Mayert", Age: 31, Gender: "Male", Address: ""},
		{ID: 4, Name: "Prof. Felicia Sauer", Age: 68, Gender: "Male", Address: ""},
		{ID: 5, Name: "Mrs. Nyah Daugherty", Age: 63, Gender: "Female", Address: ""},
		{ID: 6, Name: "Princess Hulda Glover", Age: 36, Gender: "Female", Address: ""},
		{ID: 7, Name: "Prof. Aglae Beer", Age: 77, Gender: "Female", Address: ""},
		{ID: 8, Name: "Lady Kacie Toy", Age: 69, Gender: "Male", Address: ""},
		{ID: 9, Name: "Queen Evelyn Rutherford", Age: 93, Gender: "Male", Address: ""},
		{ID: 10, Name: "Mrs. Alessandra Moore", Age: 32, Gender: "Male", Address: ""},
		{ID: 11, Name: "", Age: 0, Gender: "Male", Address: "123 Main St, City, Country"},
	}

	// Set response header dan kirim response JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patients)
}

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	config.DB.Create(&patient)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

// ==================== DOCTOR ====================

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	var doctors []models.Doctor
	config.DB.Find(&doctors)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctors)
}

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor
	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	config.DB.Create(&doctor)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctor)
}

// ==================== APPOINTMENT ====================

func GetAppointments(w http.ResponseWriter, r *http.Request) {
	var apps []models.Appointment
	config.DB.Preload("Patient").Preload("Doctor").Find(&apps)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apps)
}

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var app models.Appointment
	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	config.DB.Create(&app)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(app)
}
