package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/application/user"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/responses"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/utils"
)

func Find(ctx *gin.Context) {

	var (
		req    map[string]interface{}
		params *user.Params
		err    error
	)

	_, existNM := ctx.GetQuery("nome")
	_, existWT := ctx.GetQuery("whatsapp")
	_, existID := ctx.GetQuery("id")

	if existID || existNM || existWT {
		if err = ctx.ShouldBindQuery(&params); err != nil {
			responses.HandleError(ctx, http.StatusBadRequest, err, nil)
			return
		}
		req = utils.StructToMap(params)
	}

	response, err := user.Find(ctx, req)
	if err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, response)
}

func Create(ctx *gin.Context) {

	var (
		req *user.UserReq
		err error
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err = user.Create(ctx, req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil)
}

func Update(ctx *gin.Context) {
	var (
		req *user.UserReq
		err error
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err = user.Update(ctx, req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil)
}

func Delete(ctx *gin.Context) {
	type CampoObrigatorio struct {
		Id *string `form:"id" binding:"required"`
	}

	if err := ctx.ShouldBindQuery(new(CampoObrigatorio)); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	if err := user.Delete(ctx, new(CampoObrigatorio).Id); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil)
}
