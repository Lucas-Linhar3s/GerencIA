package users

import "github.com/Lucas-Linhar3s/GerencIA/backend/server/core/dtos"

type IUsersUsecase interface {
	Register(user *dtos.UserDTO) error
	Login(user *dtos.UserDTO) (string, error)
}
