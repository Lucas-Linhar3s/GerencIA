package user

import (
	"time"

	"github.com/google/uuid"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/utils"
)

// Service is the user service
type Service struct {
	repo IUser
}

// NewService returns a new user service
func GetService(user IUser) *Service {
	return &Service{
		repo: user,
	}
}
func GetRepository(db *database.Database) IUser {
	return newRepository(db)
}

// Find finds users
func (s *Service) Find(params map[string]interface{}) ([]UserModel, error) {
	res, err := s.repo.Find(params)
	if err != nil {
		return nil, err
	}

	dados := make([]UserModel, len(res))
	for i := 0; i < len(res); i++ {
		dados[i] = UserModel{
			ID:        res[i].ID,
			Name:      res[i].Name,
			Whatsapp:  res[i].Whatsapp,
			CreatedAt: res[i].CreatedAt,
			UpdatedAt: res[i].UpdatedAt,
		}
	}

	return dados, nil
}

// Create creates a new user
func (s *Service) Create(req *UserModel) error {
	user := UserModel{
		ID:        utils.ToStringPointer(uuid.New().String()),
		Name:      req.Name,
		Whatsapp:  req.Whatsapp,
		CreatedAt: utils.ToTimePointer(time.Now()),
	}

	if err := s.repo.Create(&user); err != nil {
		return err
	}

	return nil
}

// Update updates a user
func (s *Service) Update(req *UserModel) error {
	user := UserModel{
		ID:        req.ID,
		Name:      req.Name,
		Whatsapp:  req.Whatsapp,
		UpdatedAt: utils.ToTimePointer(time.Now()),
	}

	if err := s.repo.Update(&user); err != nil {
		return err
	}

	return nil
}

// Delete deletes a user
func (s *Service) Delete(id *string) error {
	return s.repo.Delete(id)
}
