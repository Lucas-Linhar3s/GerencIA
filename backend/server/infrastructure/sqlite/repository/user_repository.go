package repository

import (
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/Lucas-Linhar3s/GerencIA/backend/pkg/jwt"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/core/domain/users"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
)

type userRepository struct {
	db  *database.Database
	jwt *jwt.JWT
}

func NewUserRepository(db *database.Database, jwt *jwt.JWT) users.IUsersRepository {
	return &userRepository{
		db:  db,
		jwt: jwt,
	}
}

// Login implements users.IUsersRepository.
func (u *userRepository) Login(user *users.UserModel) (string, error) {
	var existe *bool
	if err := u.db.Builder.
		Select("id", "COUNT(id) > 0").
		From("pessoas").
		Where(sq.Eq{"whatsapp": user.Whatsapp}).
		QueryRow().
		Scan(&user.ID, &existe); err != nil {
		return "", err
	}

	if !*existe {
		return "", errors.New("user not found")
	}

	token, err := u.jwt.GenToken(string(user.ID.String()), time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}
	return token, err
}

// Register implements users.IUsersRepository.
func (u *userRepository) Register(user *users.UserModel) error {
	var existe *bool
	if err := u.db.Builder.
		Select("COUNT(id) > 0").
		From("pessoas").
		Where(sq.Eq{"whatsapp": user.Whatsapp}).
		QueryRow().
		Scan(&existe); err != nil {
		return err
	}

	if *existe {
		return errors.New("user already exists")
	}

	if _, err := u.db.Builder.
		Insert("pessoas").
		Columns("id", "whatsapp", "nome").
		Values(user.ID, user.Whatsapp, user.Name).
		Exec(); err != nil {
		return err
	}

	return nil
}
