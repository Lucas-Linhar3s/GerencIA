package repository

import (
	"errors"

	sq "github.com/Masterminds/squirrel"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/domain/cards"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
)

type cardRepository struct {
	db *database.Database
}

func NewCardRepository(db *database.Database) cards.ICardsRepository {
	return &cardRepository{
		db: db,
	}
}

// Delete implements cards.ICardsRepository.
func (c *cardRepository) Delete(id string) error {
	var existe *bool
	if err := c.db.Builder.
		Select("COUNT(id) > 0").
		From("cartoes").
		Where(sq.Eq{"id": id}).
		QueryRow().
		Scan(&existe); err != nil {
		return err
	}

	if !*existe {
		return errors.New("card not found")
	}

	if _, err := c.db.Builder.
		Delete("cartoes").
		Where(sq.Eq{"id": id}).
		Exec(); err != nil {
		return err
	}

	return nil
}

// Find implements cards.ICardsRepository.
func (c *cardRepository) Find(offset *int, limit *int, search *string) ([]cards.CardModel, error) {
	rows, err := c.db.Builder.
		Select("id", "nome", "limite", "created_at", "updated_at").
		From("cartoes").
		OrderBy("id").
		Limit(uint64(*limit)).
		Offset(uint64(*offset)).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []cards.CardModel
	for rows.Next() {
		var card cards.CardModel
		if err := rows.Scan(
			&card.ID,
			&card.Name,
			&card.Limit,
			&card.CreateAt,
			&card.UpdateAt,
		); err != nil {
			return nil, err
		}
		res = append(res, card)
	}
	return res, nil
}

// Update implements cards.ICardsRepository.
func (c *cardRepository) Update(card *cards.CardModel) error {
	if _, err := c.db.Builder.
		Update("cartoes").
		Set("nome", card.Name).
		Set("limite", card.Limit).
		Where(sq.Eq{"id": card.ID}).
		Exec(); err != nil {
		return err
	}

	return nil
}

func (c *cardRepository) Create(card *cards.CardModel) error {
	if _, err := c.db.Builder.
		Insert("cartoes").
		Columns(
			"id",
			"nome",
			"limite",
			"created_at",
		).
		Values(
			card.ID,
			card.Name,
			card.Limit,
			card.CreateAt,
		).
		Exec(); err != nil {
		return err
	}
	return nil
}
