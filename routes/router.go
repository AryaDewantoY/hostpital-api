package routes

import (
	"hospital-api/handlers"
	"net/http"

	_ "hospital-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Rute login
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")

	// Rute untuk mendapatkan daftar pasien dan membuat pasien baru
	r.Handle("/api/patients", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetPatients))).Methods("GET")
	r.Handle("/api/patients", handlers.AuthMiddleware(http.HandlerFunc(handlers.CreatePatient))).Methods("POST")

	// Rute untuk dokter dan janji temu
	r.Handle("/api/doctors", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetDoctors))).Methods("GET")
	r.HandleFunc("/api/appointments", handlers.CreateAppointment).Methods("POST")
	r.HandleFunc("/api/appointments", handlers.GetAppointments).Methods("GET")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.Handle("/api/patients", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetPatients))).Methods("GET")

	return r
}
