package card

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/application/card"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/responses"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/utils"
)

func Create(ctx *gin.Context) {
	var (
		req *card.CardReq
		err error
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err = card.Create(ctx, req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil)
}

func Find(ctx *gin.Context) {
	var (
		req    map[string]interface{}
		params *card.ParamsCard
		err    error
	)

	_, existNM := ctx.GetQuery("nome")
	_, existID := ctx.GetQuery("id")

	if existID || existNM {
		if err = ctx.ShouldBindQuery(&params); err != nil {
			responses.HandleError(ctx, http.StatusBadRequest, err, nil)
			return
		}
		req = utils.StructToMap(params)
	}

	response, err := card.Find(ctx, req)
	if err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, response)
}

func Update(ctx *gin.Context) {
	var (
		req *card.CardReq
		err error
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err = card.Update(ctx, req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil)
}

func Delete(ctx *gin.Context) {
	type CampoObrigatorio struct {
		Id *string `form:"id" binding:"required"`
	}

	var (
		req *CampoObrigatorio
		err error
	)

	if err = ctx.ShouldBindQuery(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err = card.Delete(ctx, req.Id); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil)
}
