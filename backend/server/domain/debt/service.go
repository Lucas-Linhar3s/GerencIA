package debt

import (
	"time"

	"github.com/google/uuid"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/utils"
)

type Service struct {
	repo IDebt
}

func GetService(debt IDebt) *Service {
	return &Service{
		repo: debt,
	}
}

func GetRepository(db *database.Database) IDebt {
	return newRepository(db)
}

func (s *Service) Find(params map[string]interface{}) ([]DebtModel, error) {
	res, err := s.repo.Find(params)
	if err != nil {
		return nil, err
	}

	resp := make([]DebtModel, len(res))
	for r := range res {
		resp[r] = DebtModel{
			ID:          res[r].ID,
			Description: res[r].Description,
			UserID:      res[r].UserID,
			CardID:      res[r].CardID,
			ValueDebt:   res[r].ValueDebt,
			Parcels:     res[r].Parcels,
			IsPaid:      res[r].IsPaid,
			CreatedAt:   res[r].CreatedAt,
			UpdatedAt:   res[r].UpdatedAt,
			UserName:    res[r].UserName,
			CardName:    res[r].CardName,
			Total:       res[r].Total,
		}
	}

	return resp, nil
}

func (s *Service) Create(req *DebtModel) error {
	debt := DebtModel{
		ID:          utils.ToStringPointer(uuid.New().String()),
		Description: req.Description,
		UserID:      req.UserID,
		CardID:      req.CardID,
		ValueDebt:   req.ValueDebt,
		Parcels:     req.Parcels,
		IsPaid:      utils.ToBoolPointer(false),
		CreatedAt:   utils.ToTimePointer(time.Now()),
	}
	if err := s.repo.Create(&debt); err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(req *DebtModel) error {
	debt := DebtModel{
		ID:          req.ID,
		Description: req.Description,
		UserID:      req.UserID,
		CardID:      req.CardID,
		ValueDebt:   req.ValueDebt,
		Parcels:     req.Parcels,
		IsPaid:      req.IsPaid,
		UpdatedAt:   utils.ToTimePointer(time.Now()),
	}
	if err := s.repo.Update(&debt); err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
