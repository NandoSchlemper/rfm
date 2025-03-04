package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	m "rfmtransportes/internal/middlewares"
	"rfmtransportes/internal/models"
	"rfmtransportes/internal/services"
)

func RegisterUserHandlers(
	mux *http.ServeMux,
) {
	mux.Handle("POST user/create", m.NewSetDBContext(handleCreateUser))
	mux.Handle("POST auth/login", m.NewSetDBContext(handleLoginUser))
	mux.Handle("POST auth/register", m.NewSetDBContext(handleRegisterUser))
	mux.Handle("GET auth/test", m.SecureRoute(handleToTestUser))
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	userService := services.UserService{DB: db}
	var userBody *models.UserDTO
	decodeError := json.NewDecoder(r.Body).Decode(&userBody)
	if decodeError != nil {
		fmt.Printf("Erro ao solicitar o payload\n%s", decodeError.Error())
		http.Error(w, decodeError.Error(), http.StatusBadRequest)
		return
	}

	createError := userService.Create(userBody)
	if createError != nil {
		fmt.Printf("Erro ao utilizar o userService para criar o usuário\n%s", createError.Error())
		http.Error(w, createError.Error(), http.StatusBadRequest)
		return
	}
}

func handleLoginUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// payload, validar, gerar token, meter no Header
	userService := &services.AuthService{DB: db}
	var userBody *models.User
	decodeError := json.NewDecoder(r.Body).Decode(&userBody)

	if decodeError != nil {
		fmt.Printf("Erro ao solicitar payload de login\n%s", decodeError.Error())
		http.Error(w, decodeError.Error(), http.StatusBadRequest)
		return
	}

	token, err := userService.Login(userBody)
	if err != nil {
		fmt.Printf("Erro ao realizar o login do usuário\n%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	fmt.Fprint(w, "Login bem sucedido")
}

func handleRegisterUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	authService := &services.AuthService{DB: db}
	var userBody *models.User
	decodeError := json.NewDecoder(r.Body).Decode(&userBody)

	if decodeError != nil {
		fmt.Printf("Erro ao solicitar payload de registro\n%s", decodeError.Error())
		http.Error(w, decodeError.Error(), http.StatusBadRequest)
		return
	}

	token, err := authService.Register(userBody)
	if err != nil {
		fmt.Printf("Erro ao registrar o usuário\n%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	fmt.Fprint(w, "Registro bem sucedido")
}

func handleToTestUser(w http.ResponseWriter, r *http.Request, user *models.User) {
	r.Header.Get("token")
	fmt.Fprint(w, "Testando")
}
