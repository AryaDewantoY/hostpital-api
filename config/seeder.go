package config

import (
	"fmt"
	"hospital-api/models"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Fungsi untuk memilih acak dari pilihan
func randomChoice(options []string) string {
	rand.Seed(time.Now().UnixNano())
	return options[rand.Intn(len(options))]
}

func SeedUsers(db *gorm.DB) {
	// Contoh isi fungsi seeding
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	db.Create(&user) // Menyimpan data user ke database
}

// SeedPatients untuk menambahkan data Patient palsu
func SeedPatients(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		patient := models.Patient{
			Name:   faker.Name(),
			Age:    rand.Intn(100),                           // Menggunakan rand.Intn untuk angka acak
			Gender: randomChoice([]string{"Male", "Female"}), // Menggunakan randomChoice untuk memilih jenis kelamin
		}

		if err := db.Create(&patient).Error; err != nil {
			fmt.Printf("failed to seed patients: %v\n", err)
		}
	}
}

// SeedDoctors untuk menambahkan data Doctor palsu
func SeedDoctors(db *gorm.DB) {
	for i := 0; i < 5; i++ {
		doctor := models.Doctor{
			Name:      faker.Name(),
			Specialty: randomChoice([]string{"Cardiology", "Neurology", "Orthopedics", "Pediatrics"}), // Menggunakan randomChoice
		}

		if err := db.Create(&doctor).Error; err != nil {
			fmt.Printf("failed to seed doctors: %v\n", err)
		}
	}
}

// SeedAppointments untuk menambahkan data Appointment palsu
func SeedAppointments(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		appointment := models.Appointment{
			PatientID: uint(rand.Intn(10) + 1),
			DoctorID:  uint(rand.Intn(5) + 1),
		}

		if err := db.Create(&appointment).Error; err != nil {
			fmt.Printf("failed to seed appointments: %v\n", err)
		}
	}
}

// Seeder untuk semua data
func SeedAll(db *gorm.DB) {
	SeedUsers(db)
	SeedPatients(db)
	SeedDoctors(db)
	SeedAppointments(db)
}
