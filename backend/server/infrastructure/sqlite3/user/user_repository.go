package user

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
)

type UserRepository struct {
	DB *database.Database
}

func (r *UserRepository) Create(req *UserModel) error {
	if err := r.DB.Builder.
		Insert("usuarios").
		Columns("id", "nome", "whatsapp", "created_at").
		Values(req.ID, req.Name, req.Whatsapp, req.CreatedAt).
		Suffix("RETURNING id").
		Scan(new(string)); err != nil {
		return err
	}
	return nil
}

// Find finds users
func (r *UserRepository) Find(params map[string]interface{}) ([]UserModel, error) {
	consulta := r.DB.Builder.
		Select("id", "nome", "whatsapp", "created_at", "updated_at").
		From("usuarios")

	if params != nil {
		consulta = consulta.Where(sq.Or{
			sq.Eq{"id": params["Id"]},
			sq.Eq{"nome": params["Nome"]},
			sq.Eq{"whatsapp": params["Whatsapp"]},
		})
	}

	rows, err := consulta.Query()
	if err != nil {
		return nil, err
	}

	var users []UserModel
	for rows.Next() {
		var user UserModel
		if err := rows.Scan(&user.ID, &user.Name, &user.Whatsapp, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Update a user by id
func (r *UserRepository) Update(req *UserModel) error {
	_, err := r.DB.Builder.
		Update("usuarios").
		SetMap(sq.Eq{"nome": req.Name, "whatsapp": req.Whatsapp, "updated_at": req.UpdatedAt}).
		Where(sq.Eq{"id": req.ID}).
		Exec()
	return err
}

// Delete a user by id
func (r *UserRepository) Delete(id *string) error {
	_, err := r.DB.Builder.
		Delete("usuarios").
		Where(sq.Eq{"id": id}).
		Exec()
	return err
}
