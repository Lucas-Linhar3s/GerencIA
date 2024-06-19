package usecases

import (
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/domain/cards"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/dtos"
)

type cardUsecaseImpl struct {
	repo cards.ICardsRepository
}

// NewCardUsecase implements cards.ICardUsecase.
func NewCardUsecase(repo cards.ICardsRepository) cards.ICardUsecase {
	return &cardUsecaseImpl{repo: repo}
}

// Create implements cards.ICardUsecase.
func (c *cardUsecaseImpl) Create(card *dtos.CardDTO) error {
	cardModel := cards.NewCard(
		card.Name,
		card.Limit,
		true)

	if err := cardModel.Validate(); err != nil {
		return err
	}

	return c.repo.Create(cardModel)
}

// Delete implements cards.ICardUsecase.
func (c *cardUsecaseImpl) Delete(id string) error {
	return c.repo.Delete(id)
}

// Find implements cards.ICardUsecase.
func (c *cardUsecaseImpl) Find(params *dtos.ParamsDTO) (res []dtos.CardDTO, err error) {
	data, err := c.repo.Find(
		params.Offset,
		params.Limit,
		params.Search,
	)

	res = make([]dtos.CardDTO, len(data))
	for i := range data {
		res[i] = dtos.CardDTO{
			ID:       data[i].ID,
			Name:     data[i].Name,
			Limit:    data[i].Limit,
			CreateAt: data[i].CreateAt,
			UpdateAt: data[i].UpdateAt,
		}
	}
	return
}

// Update implements cards.ICardUsecase.
func (c *cardUsecaseImpl) Update(card *dtos.CardDTO) error {
	cardModel := cards.CardModel{
		ID:    card.ID,
		Name:  card.Name,
		Limit: card.Limit,
	}

	if err := cardModel.Validate(); err != nil {
		return err
	}

	if err := c.repo.Update(&cardModel); err != nil {
		return err
	}
	return nil
}
