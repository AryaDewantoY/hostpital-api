package handlers

import (
	"hospital-api/config"
	"hospital-api/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Fungsi untuk menghasilkan JWT
func generateJWT(user models.User) (string, error) {
	// Membuat klaim
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token kadaluarsa dalam 72 jam
	}

	// Membuat token JWT dengan menggunakan klaim dan secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token menggunakan Secret yang dideklarasikan di config.go
	signedToken, err := token.SignedString(config.Secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
