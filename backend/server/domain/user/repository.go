package user

import (
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/infrastructure/sqlite3/user"
)

type Repository struct {
	repo *user.UserRepository
}

func NewRepository(db *database.Database) IUser {
	return &Repository{
		repo: &user.UserRepository{
			DB: db,
		},
	}
}

// Create implements IUser.
func (r *Repository) Create(req *UserModel) error {
	user := user.UserModel{
		ID:        req.ID,
		Name:      req.Name,
		Whatsapp:  req.Whatsapp,
		CreatedAt: req.CreatedAt,
	}
	if err := r.repo.Create(&user); err != nil {
		return err
	}

	return nil
}

// Delete implements IUser.
func (r *Repository) Delete(id *string) error {
	return r.repo.Delete(id)
}

// Find implements IUser.
func (r *Repository) Find(params map[string]interface{}) ([]UserModel, error) {
	res, err := r.repo.Find(params)
	if err != nil {
		return nil, err
	}

	var resp = make([]UserModel, len(res))
	for r := range res {
		resp[r] = UserModel{
			ID:        res[r].ID,
			Name:      res[r].Name,
			Whatsapp:  res[r].Whatsapp,
			CreatedAt: res[r].CreatedAt,
			UpdatedAt: res[r].UpdatedAt,
		}
	}

	return resp, nil
}

// Update implements IUser.
func (r *Repository) Update(req *UserModel) error {
	user := user.UserModel{
		ID:        req.ID,
		Name:      req.Name,
		Whatsapp:  req.Whatsapp,
		UpdatedAt: req.UpdatedAt,
	}

	if err := r.repo.Update(&user); err != nil {
		return err
	}

	return nil
}
