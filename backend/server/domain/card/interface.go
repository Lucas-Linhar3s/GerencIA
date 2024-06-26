package card

type ICard interface {
	Find(params map[string]interface{}) ([]CardModel, error)
	Create(req *CardModel) error
	Update(req *CardModel) error
	Delete(id *string) error
}
