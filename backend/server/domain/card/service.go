package card

import (
	"time"

	"github.com/google/uuid"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/utils"
)

// Service is the card service
type Service struct {
	repo ICard
}

// NewService returns a new card service
func GetService(card ICard) *Service {
	return &Service{
		repo: card,
	}
}

// GetRepository returns a new card repository
func GetRepository(db *database.Database) ICard {
	return newRepository(db)
}

// Find finds cards
func (s *Service) Find(params map[string]interface{}) ([]CardModel, error) {
	res, err := s.repo.Find(params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create creates a new card
func (s *Service) Create(req *CardModel) error {
	card := CardModel{
		ID:      utils.ToStringPointer(uuid.New().String()),
		Nome:    req.Nome,
		Limite:  req.Limite,
		Created: utils.ToTimePointer(time.Now()),
	}

	if err := s.repo.Create(&card); err != nil {
		return err
	}

	return nil
}

// Update updates a card
func (s *Service) Update(req *CardModel) error {
	card := CardModel{
		ID:      req.ID,
		Nome:    req.Nome,
		Limite:  req.Limite,
		Updated: utils.ToTimePointer(time.Now()),
	}

	if err := s.repo.Update(&card); err != nil {
		return err
	}

	return nil
}

// Delete deletes a card
func (s *Service) Delete(id *string) error {
	return s.repo.Delete(id)
}
