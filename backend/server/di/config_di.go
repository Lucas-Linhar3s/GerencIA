package di

import (
	"log"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/handler"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/service"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/sqlite/repository"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/usecases"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
)

var servicesDI service.ICardServices

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
