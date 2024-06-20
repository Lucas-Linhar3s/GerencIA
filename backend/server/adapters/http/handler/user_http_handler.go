package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/responses"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/service"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/domain/users"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/dtos"
)

type UserHandler struct {
	service users.IUsersUsecase
}

func NewUserHandler(service users.IUsersUsecase) service.IUserService {
	return &UserHandler{
		service: service,
	}
}

// Login implements service.IUserService.
func (u *UserHandler) Login(ctx *gin.Context) {
	var req dtos.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	token, err := u.service.Login(&req)
	if err != nil {
		responses.HandleError(ctx, http.StatusUnauthorized, err, nil)
		return
	}
	responses.HandleSuccess(ctx, token, http.StatusOK)
}

// Register implements service.IUserService.
func (u *UserHandler) Register(ctx *gin.Context) {
	var req dtos.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	err := u.service.Register(&req)
	if err != nil {
		responses.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	responses.HandleSuccess(ctx, nil, http.StatusNoContent)
}
