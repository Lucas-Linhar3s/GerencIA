package usecases

import (
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/domain/users"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/dtos"
)

type UserUsecaseImpl struct {
	repo users.IUsersRepository
}

func NewUserUsecase(repo users.IUsersRepository) users.IUsersUsecase {
	return &UserUsecaseImpl{repo: repo}
}

// Login implements users.IUsersUsecase.
func (u *UserUsecaseImpl) Login(user *dtos.UserDTO) (string, error) {
	userModel := users.NewUser(nil, user.Whatsapp, true)
	if err := userModel.Validate(); err != nil {
		return "", err
	}

	return u.repo.Login(userModel)
}

// Register implements users.IUsersUsecase.
func (u *UserUsecaseImpl) Register(user *dtos.UserDTO) error {
	userModel := users.NewUser(user.Name, user.Whatsapp, true)
	if err := userModel.Validate(); err != nil {
		return err
	}

	return u.repo.Register(userModel)
}
