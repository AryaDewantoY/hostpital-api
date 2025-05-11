package controllers

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// implementasi login
	json.NewEncoder(w).Encode(map[string]string{"message": "Login success"})
}
