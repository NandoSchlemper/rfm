package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"rfmtransportes/internal/database"
	"rfmtransportes/internal/models"
	"rfmtransportes/internal/services"
)

func RegisterDriverHandlers(
	mux *http.ServeMux,
) {
	mux.HandleFunc("GET /", handleNothing)
	mux.HandleFunc("POST /drivers/create", handleCreateUser)
}

func handleNothing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! this is my application"))
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	driverService := &services.DriverService{}

	var dto models.DriverDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	creating := driverService.Create(dto, &database.NeonDB{})
	if creating != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
