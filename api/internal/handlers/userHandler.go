package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"rfmtransportes/internal/middlewares"
	"rfmtransportes/internal/models"
	"rfmtransportes/internal/services"
)

func RegisterUserHandlers(
	mux *http.ServeMux,
) {
	mux.Handle("POST user/create", middlewares.NewSetDBContext(handleCreateUser))
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	userService := services.UserService{DB: db}
	var userBody *models.UserDTO
	request := json.NewDecoder(r.Body).Decode(&userBody)

	if request.Error != nil {
		fmt.Printf("Erro ao pegar o payload,\n%s", request.Error)
		return
	}

	err := userService.Create(userBody)
	if err != nil {
		fmt.Printf("Erro ao utilizar o userService para criar o usuário\n%s", err.Error)
	}
}
