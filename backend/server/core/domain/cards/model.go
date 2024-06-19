package cards

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type CardModel struct {
	ID       uuid.UUID
	Name     *string
	Limit    *float64
	CreateAt *time.Time
	UpdateAt *time.Time
}

func NewCard(name *string, limit *float64, create bool) *CardModel {
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

	return &CardModel{
		ID:       uuid.New(),
		Name:     name,
		Limit:    limit,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}
}

func (c *CardModel) Validate() error {
	if c.Name == nil || *c.Name == "" {
		return errors.New("invalid card name")
	} else if c.Limit == nil || *c.Limit < 0 {
		return errors.New("invalid card limit")
	}
	return nil
}
