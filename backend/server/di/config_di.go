package di

import (
	"log"

	"github.com/Lucas-Linhar3s/GerencIA/backend/pkg/jwt"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/handler"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/service"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/usecases"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/infrastructure/sqlite/repository"
)

var servicesDI service.ICardServices
var userDI service.IUserService

func ConfigDi(db *database.Database) service.ICardServices {
	cardRepository := repository.NewCardRepository(db)
	cardUsecase := usecases.NewCardUsecase(cardRepository)
	servicesDI = handler.NewCardHandler(cardUsecase)

	return servicesDI
}

// GetConfig returns a pointer to a Config struct which holds a valid config
func GetServices() service.ICardServices {
	if servicesDI == nil {
		log.Fatal("config was not successfully loaded")
	}
	return servicesDI
}

func UserDi(db *database.Database, jwt *jwt.JWT) service.IUserService {
	userRepository := repository.NewUserRepository(db, jwt)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userDI = handler.NewUserHandler(userUsecase)
	return userDI
}

func GetUserService() service.IUserService {
	if servicesDI == nil {
		log.Fatal("config was not successfully loaded")
	}
	return userDI
}
