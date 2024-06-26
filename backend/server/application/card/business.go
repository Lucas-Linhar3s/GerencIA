package card

import (
	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/di"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/domain/card"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/responses"
)

// Create creates a card
func Create(ctx *gin.Context, req *CardReq) error {
	const msg = "Error creating card"

	var (
		service = card.GetService(card.GetRepository(di.GetDatabase()))
	)

	card := card.CardModel{
		Nome:   req.Nome,
		Limite: req.Limite,
	}

	if err := service.Create(&card); err != nil {
		responses.CheckError(msg, err)
	}

	return nil
}

// Update updates a card
func Update(ctx *gin.Context, req *CardReq) error {
	const msg = "Error updating card"

	var (
		service = card.GetService(card.GetRepository(di.GetDatabase()))
	)

	card := card.CardModel{
		ID:     req.Id,
		Nome:   req.Nome,
		Limite: req.Limite,
	}

	if err := service.Update(&card); err != nil {
		responses.CheckError(msg, err)
	}

	return nil
}

// Delete deletes a card
func Delete(ctx *gin.Context, id *string) error {
	const msg = "Error deleting card"

	var (
		service = card.GetService(card.GetRepository(di.GetDatabase()))
	)

	if err := service.Delete(id); err != nil {
		responses.CheckError(msg, err)
	}

	return nil
}

// Find finds cards
func Find(ctx *gin.Context, params map[string]interface{}) (*[]CardRes, error) {
	const msg = "Error finding cards"

	var (
		service = card.GetService(card.GetRepository(di.GetDatabase()))
	)

	res, err := service.Find(params)
	if err != nil {
		responses.CheckError(msg, err)
	}

	if len(res) == 0 {
		return nil, responses.ErrNotFound
	}

	var resp = make([]CardRes, len(res))
	for v := range res {
		resp[v] = CardRes{
			ID:      res[v].ID,
			Nome:    res[v].Nome,
			Limite:  res[v].Limite,
			Created: res[v].Created,
			Updated: res[v].Updated,
		}
	}

	return &resp, nil
}
