package users

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID       uuid.UUID
	Name     *string
	Whatsapp *string
	CreateAt *time.Time
	UpdateAt *time.Time
}

func NewUser(name *string, whatsapp *string, create bool) *UserModel {
	var (
		now      = time.Now()
		createAt *time.Time
		updateAt *time.Time
	)

	if create {
		createAt = &now
	} else {
		updateAt = &now
	}

	return &UserModel{
		ID:       uuid.New(),
		Name:     name,
		Whatsapp: whatsapp,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}
}

func (c *UserModel) Validate() error {
	if c.Whatsapp == nil {
		return errors.New("invalid number whatsapp")
	}
	return nil
}
