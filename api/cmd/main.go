package main

import (
	"fmt"

	s "rfmtransportes/internal/server"
)

func main() {
	// db := d.NeonDB{}
	// db.Migrate()
	server := s.NewServer()
	err := server.Start()
	if err != nil {
		fmt.Printf("Erro ao instanciaro server%v\n", err.Error())
	}
}
