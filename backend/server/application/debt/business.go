package debt

import (
	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/di"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/domain/debt"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/responses"
)

func Find(ctx *gin.Context, params map[string]interface{}) ([]Res, error) {
	const msg = "Error finding debt"

	var (
		service = debt.GetService(debt.GetRepository(di.GetDatabase()))
	)

	debt, err := service.Find(params)
	if err != nil {
		return nil, responses.CheckError(msg, err)
	}

	if len(debt) == 0 {
		return nil, responses.ErrNotFound
	}

	resp := make([]Res, len(debt))
	for r := range debt {
		resp[r] = Res{
			ID:          debt[r].ID,
			Description: debt[r].Description,
			ValueDebt:   debt[r].ValueDebt,
			Parcels:     debt[r].Parcels,
			IsPaid:      debt[r].IsPaid,
			CreatedAt:   debt[r].CreatedAt,
			UpdatedAt:   debt[r].UpdatedAt,
			Total:       debt[r].Total,
			User: User{
				ID:   debt[r].UserID,
				Name: debt[r].UserName,
			},
			Card: Card{
				ID:   debt[r].CardID,
				Name: debt[r].CardName,
			},
		}
	}

	return resp, nil
}

func Create(ctx *gin.Context, req Req) error {
	const msg = "Error creating debt"
	var (
		service = debt.GetService(debt.GetRepository(di.GetDatabase()))
	)

	debt := debt.DebtModel{
		Description: req.Description,
		UserID:      req.UserID,
		CardID:      req.CardID,
		ValueDebt:   req.ValueDebt,
		Parcels:     req.Parcels,
	}

	if err := service.Create(&debt); err != nil {
		return responses.CheckError(msg, err)
	}

	return nil
}

func Update(ctx *gin.Context, req Req) error {
	const msg = "Error updating debt"
	var (
		service = debt.GetService(debt.GetRepository(di.GetDatabase()))
	)

	debt := debt.DebtModel{
		ID:          req.ID,
		Description: req.Description,
		UserID:      req.UserID,
		CardID:      req.CardID,
		ValueDebt:   req.ValueDebt,
		Parcels:     req.Parcels,
		IsPaid:      req.IsPaid,
	}

	if err := service.Update(&debt); err != nil {
		return responses.CheckError(msg, err)
	}

	return nil
}

func Delete(ctx *gin.Context, id *string) error {
	const msg = "Error deleting debt"
	var (
		service = debt.GetService(debt.GetRepository(di.GetDatabase()))
	)

	if err := service.Delete(id); err != nil {
		return responses.CheckError(msg, err)
	}

	return nil
}
