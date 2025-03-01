package handlers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"

	"rfmtransportes/internal/middlewares"
	"rfmtransportes/internal/models"
	"rfmtransportes/internal/services"
)

func RegisterDriverHandlers(
	mux *http.ServeMux,
) {
	mux.HandleFunc("GET /", handleNothing)
	mux.Handle("POST /drivers/create", middlewares.NewSetDBContext(handleCreateUser))
}

func handleNothing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! this is my application"))
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	driverService := &services.DriverService{DB: db}

	var dto models.DriverDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbInsertError := driverService.CreateDriver(dto)

	if dbInsertError != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
