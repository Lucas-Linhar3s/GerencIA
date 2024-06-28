package debt

import "time"

type DebtModel struct {
	ID          *string
	Description *string
	UserID      *string
	UserName    *string
	CardID      *string
	CardName    *string
	ValueDebt   *float64
	Parcels     *string
	IsPaid      *bool
	Total       *float64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
