package card

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
)

type SQCardRepository struct {
	Db *database.Database
}

func (r *SQCardRepository) Create(req *CardModel) error {
	if err := r.Db.Builder.
		Insert("cartoes").
		Columns("id", "nome", "limite", "created_at").
		Values(req.ID, req.Nome, req.Limite, req.Created).
		Suffix("RETURNING id").
		Scan(new(string)); err != nil {
		return err
	}
	return nil
}

func (r *SQCardRepository) Find(params map[string]interface{}) ([]CardModel, error) {
	consulta := r.Db.Builder.
		Select("id", "nome", "limite", "created_at", "updated_at").
		From("cartoes")

	if params != nil {
		p := params["Nome"].(*string)
		search := "%" + *p + "%"
		consulta = consulta.Where(sq.Or{
			sq.Expr("id = ? COLLATE NOCASE", params["Id"]),
			sq.Expr("nome LIKE ? COLLATE NOCASE", search),
		})
	}

	rows, err := consulta.Query()
	if err != nil {
		return nil, err
	}

	var cards []CardModel
	for rows.Next() {
		var card CardModel
		if err := rows.Scan(&card.ID, &card.Nome, &card.Limite, &card.Created, &card.Updated); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (r *SQCardRepository) Update(req *CardModel) error {
	if _, err := r.Db.Builder.
		Update("cartoes").
		SetMap(sq.Eq{"nome": req.Nome, "limite": req.Limite, "updated_at": req.Updated}).
		Where(sq.Eq{"id": req.ID}).
		Exec(); err != nil {
		return err
	}
	return nil
}

func (r *SQCardRepository) Delete(id *string) error {
	if _, err := r.Db.Builder.
		Delete("cartoes").
		Where(sq.Eq{"id": id}).
		Exec(); err != nil {
		return err
	}
	return nil
}
