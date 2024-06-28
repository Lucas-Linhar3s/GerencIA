package debt

import (
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/infrastructure/sqlite3/debt"
)

type repository struct {
	sqDebt *debt.SQDebt
}

func newRepository(db *database.Database) *repository {
	return &repository{
		sqDebt: &debt.SQDebt{
			Db: db,
		},
	}
}

// Create implements IDebt.
func (r *repository) Create(req *DebtModel) error {
	debt := debt.DebtModel{
		ID:          req.ID,
		Description: req.Description,
		UserID:      req.UserID,
		CardID:      req.CardID,
		ValueDebt:   req.ValueDebt,
		Parcels:     req.Parcels,
		IsPaid:      req.IsPaid,
		CreatedAt:   req.CreatedAt,
	}

	if err := r.sqDebt.Create(&debt); err != nil {
		return err
	}

	return nil
}

// Delete implements IDebt.
func (r *repository) Delete(id *string) error {
	return r.sqDebt.Delete(id)
}

// Find implements IDebt.
func (r *repository) Find(params map[string]interface{}) ([]DebtModel, error) {
	res, err := r.sqDebt.Find(params)
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

// Update implements IDebt.
func (r *repository) Update(req *DebtModel) error {
	debt := debt.DebtModel{
		ID:          req.ID,
		Description: req.Description,
		UserID:      req.UserID,
		CardID:      req.CardID,
		ValueDebt:   req.ValueDebt,
		Parcels:     req.Parcels,
		IsPaid:      req.IsPaid,
		UpdatedAt:   req.UpdatedAt,
	}

	if err := r.sqDebt.Update(&debt); err != nil {
		return err
	}

	return nil
}
