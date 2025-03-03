package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"rfmtransportes/internal/handlers"
	"rfmtransportes/internal/services"
)

type Server struct {
	Router *http.ServeMux
	Driver *services.DriverService
	User   *services.UserService
}

func NewServer() *Server {
	s := &Server{
		Router: http.NewServeMux(),
		Driver: &services.DriverService{},
		User:   &services.UserService{},
	}
	s.addRoutes()
	return s
}

func (s *Server) addRoutes() {
	handlers.RegisterDriverHandlers(s.Router)
	handlers.RegisterUserHandlers(s.Router)
}

func (s *Server) Start() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Erro ao carregar .env")
		return errors.New("Erro ao carregar .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Print("Erro ao carregar PORT")
		return errors.New("Carregamento de port invalido")
	}
	return http.ListenAndServe(port, s.Router)
}
