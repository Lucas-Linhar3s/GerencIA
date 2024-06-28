package debt

import "time"

type Req struct {
	ID          *string  `json:"id"`
	Description *string  `json:"description"`
	UserID      *string  `json:"user_id"`
	CardID      *string  `json:"card_id"`
	ValueDebt   *float64 `json:"value_debt"`
	Parcels     *string  `json:"parcels"`
	IsPaid      *bool    `json:"is_paid"`
}

type Res struct {
	ID          *string    `json:"id"`
	Description *string    `json:"description"`
	ValueDebt   *float64   `json:"value_debt"`
	Parcels     *string    `json:"parcels"`
	IsPaid      *bool      `json:"is_paid"`
	Total       *float64   `json:"total"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	User        User       `json:"user"`
	Card        Card       `json:"card"`
}

type User struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type Card struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type Params struct {
	UserId      *string `form:"user_id"`
	CardId      *string `form:"card_id"`
	Description *string `form:"description"`
	UserName    *string `form:"user_name"`
	CardName    *string `form:"card_name"`
}
