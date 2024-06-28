package user

import "time"

// UserModel is the model for user
type UserModel struct {
	ID        *string
	Name      *string
	Whatsapp  *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
