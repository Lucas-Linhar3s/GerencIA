package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/responses"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/service"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/domain/cards"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/dtos"
)

type cardServiceImpl struct {
	service cards.ICardUsecase
}

func NewCardHandler(service cards.ICardUsecase) service.ICardServices {
	return cardServiceImpl{
		service: service,
	}
}

// Create godoc
// @Summary Create a new card
// @Description Create a new card
// @Tags cards
// @Accept json
// @Produce json
// @Param card body dtos.CardDTO true "Card Data"
// @Success 204
// @Router /card/create [post]
func (c cardServiceImpl) Create(ctx *gin.Context) {
	var card dtos.CardDTO
	if err := ctx.ShouldBindJSON(&card); err != nil {
		responses.HandleError(ctx, 400, err, nil)
		return
	}

	if err := c.service.Create(&card); err != nil {
		responses.HandleError(ctx, 400, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil, http.StatusNoContent)
}

// Delete godoc
// @Summary Delete a card by ID
// @Description Delete a card by ID
// @Tags cards
// @Accept json
// @Produce json
// @Param id path string true "Card ID"
// @Success 204 {object} responses.Response "Successfully deleted the card"
// @Router /card/delete/{id} [delete]
func (c cardServiceImpl) Delete(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	if err := uuid.Validate(id); err != nil {
		responses.HandleError(ctx, 400, err, nil)
		return
	}

	if err := c.service.Delete(id); err != nil {
		responses.HandleError(ctx, 400, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil, http.StatusNoContent)
}

// Find godoc
// @Summary Get all cards
// @Description Get all cards
// @Tags cards
// @Accept json
// @Produce json
// @Param   limit  query  int  false  "Query parameters"
// @Param   offset  query  int  false  "Query parameters"
// @Param   search  query  string  false  "Query parameters"
// @Success 200 {object} responses.Response
// @Router /card/list [get]
func (c cardServiceImpl) Find(ctx *gin.Context) {
	params, err := dtos.FromValueParamsDTO(ctx.Request)
	if err != nil {
		responses.HandleError(ctx, 400, err, nil)
		return
	}

	pagination, err := c.service.Find(params)
	if err != nil {
		responses.HandleError(ctx, 400, err, nil)
		return
	}

	if len(pagination) == 0 {
		responses.HandleSuccess(ctx, nil, http.StatusNoContent)
		return
	}

	responses.HandleSuccess(ctx, pagination, http.StatusOK)
}

// Create godoc
// @Summary Update a card
// @Description Update a card
// @Tags cards
// @Accept json
// @Produce json
// @Param card body dtos.CardDTO true "Card Data"
// @Success 204
// @Router /card/update [patch]
func (c cardServiceImpl) Update(ctx *gin.Context) {
	var card dtos.CardDTO
	if err := ctx.ShouldBindJSON(&card); err != nil {
		responses.HandleError(ctx, 400, err, nil)
		return
	}

	if err := c.service.Update(&card); err != nil {
		responses.HandleError(ctx, 400, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil, http.StatusNoContent)
}
