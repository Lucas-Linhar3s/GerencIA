package dtos

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID       uuid.UUID  `json:"id"`
	Name     *string    `json:"nome"`
	Whatsapp *string    `json:"whatsapp" binding:"required" example:"84999999999"`
	CreateAt *time.Time `json:"criado_em"`
	UpdateAt *time.Time `json:"atualizado_em"`
}

func (dto *UserDTO) FromJSON(data []byte) error {
	return json.Unmarshal(data, dto)
}

func (dto *UserDTO) ToJSON() ([]byte, error) {
	return json.Marshal(dto)
}
