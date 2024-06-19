package cards

type ICardsRepository interface {
	Create(card *CardModel) error
	Find(offset *int, limit *int, search *string) ([]CardModel, error)
	Update(card *CardModel) error
	Delete(id string) error
}
