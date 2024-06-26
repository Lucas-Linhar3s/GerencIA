package card

import (
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/infrastructure/sqlite3/card"
)

type repository struct {
	repo *card.SQCardRepository
}

func newRepository(db *database.Database) *repository {
	return &repository{
		repo: &card.SQCardRepository{
			Db: db,
		},
	}
}

// Create implements ICard.
func (r *repository) Create(req *CardModel) error {
	card := &card.CardModel{
		ID:      req.ID,
		Nome:    req.Nome,
		Limite:  req.Limite,
		Created: req.Created,
	}

	if err := r.repo.Create(card); err != nil {
		return err
	}

	return nil
}

// Delete implements ICard.
func (r *repository) Delete(id *string) error {
	return r.repo.Delete(id)
}

// Find implements ICard.
func (r *repository) Find(params map[string]interface{}) ([]CardModel, error) {
	res, err := r.repo.Find(params)
	if err != nil {
		return nil, err
	}

	var resp = make([]CardModel, len(res))
	for r := range res {
		resp[r] = CardModel{
			ID:      res[r].ID,
			Nome:    res[r].Nome,
			Limite:  res[r].Limite,
			Created: res[r].Created,
			Updated: res[r].Updated,
		}
	}

	return resp, nil
}

// Update implements ICard.
func (r *repository) Update(req *CardModel) error {
	card := card.CardModel{
		ID:      req.ID,
		Nome:    req.Nome,
		Limite:  req.Limite,
		Updated: req.Updated,
	}

	if err := r.repo.Update(&card); err != nil {
		return err
	}

	return nil
}
