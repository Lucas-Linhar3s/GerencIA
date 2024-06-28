package di

import (
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/domain/user"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/pkg/jwt"
)

var Repository user.IUser

func ConfigDI(db *database.Database, jwt *jwt.JWT) {
	Repository = user.NewRepository(db)
}
