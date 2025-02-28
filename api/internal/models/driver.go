package models

import "time"

type Driver struct {
	Id         string
	name       string
	cpf        string
	created_at time.Time
	created_by string
}

type DriverDTO struct {
	name string
	cpf  string
}
