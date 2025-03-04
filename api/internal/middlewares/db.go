package middlewares

import (
	"net/http"

	"gorm.io/gorm"

	"rfmtransportes/internal/database"
)

type HandlerWithDB func(http.ResponseWriter, *http.Request, *gorm.DB)

type setDBContext struct {
	handler HandlerWithDB
}

func (db *setDBContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	database := database.NeonDB{}
	connection := database.Connection()

	db.handler(w, r, connection)
}

func NewSetDBContext(handler HandlerWithDB) *setDBContext {
	return &setDBContext{handler}
}
