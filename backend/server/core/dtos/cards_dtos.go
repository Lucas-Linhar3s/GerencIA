package dtos

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type CardDTO struct {
	ID       uuid.UUID  `json:"id"`
	Name     *string    `json:"nome" binding:"required"`
	Limit    *float64   `json:"limite" binding:"required"`
	CreateAt *time.Time `json:"criado_em"`
	UpdateAt *time.Time `json:"atualizado_em"`
}

func (dto *CardDTO) FromJSON(data []byte) error {
	return json.Unmarshal(data, dto)
}

func (dto *CardDTO) ToJSON() ([]byte, error) {
	return json.Marshal(dto)
}
