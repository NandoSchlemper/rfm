package server

import (
	"database/sql"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
	// driverService
}

func NewServer(
	db *sql.DB,
) *Server {
	s := &Server{
		Router: http.NewServeMux(),
		// driver service
	}
	s.addRoutes()
	return s
}

func (s *Server) addRoutes() {
	// s.Router.Handle(path, handler)

	// driverHandler := handlers.NewDriverHandler(s.userService)
	// defining routes --> s.router.HandleFunc(path, driverHandler.GetDriverByID) ??
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}
