package cards

import "github.com/Lucas-Linhar3s/GerencIA/backend/server/core/dtos"

type ICardUsecase interface {
	Create(card *dtos.CardDTO) error
	Find(params *dtos.ParamsDTO) ([]dtos.CardDTO, error)
	Update(card *dtos.CardDTO) error
	Delete(id string) error
}
