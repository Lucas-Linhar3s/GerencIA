package debt

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/application/debt"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/responses"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/utils"
)

func Create(ctx *gin.Context) {
	var (
		req *debt.Req
		err error
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err = debt.Create(ctx, *req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil)
}

func Find(ctx *gin.Context) {
	var (
		req    map[string]interface{}
		params *debt.Params
		err    error
	)

	_, existNM := ctx.GetQuery("user_id")
	_, existWT := ctx.GetQuery("card_id")
	_, existID := ctx.GetQuery("description")
	_, existUN := ctx.GetQuery("user_name")
	_, existCN := ctx.GetQuery("card_name")

	if existID || existNM || existWT || existUN || existCN {
		if err = ctx.ShouldBindQuery(&params); err != nil {
			responses.HandleError(ctx, http.StatusBadRequest, err, nil)
			return
		}
		req = utils.StructToMap(params)
	}

	response, err := debt.Find(ctx, req)
	if err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, response)
}

func Update(ctx *gin.Context) {
	var (
		req *debt.Req
		err error
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err = debt.Update(ctx, *req); err != nil {
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

	if err = debt.Delete(ctx, req.Id); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil)
}
