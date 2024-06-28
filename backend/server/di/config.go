package di

import (
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/pkg/jwt"
)

var db *database.Database

func ConfigDI(conn *database.Database, jwt *jwt.JWT) {
	db = conn
}

func GetDatabase() *database.Database {
	return db
}
