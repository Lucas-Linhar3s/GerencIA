package card

import "time"

type CardModel struct {
	ID      *string
	Nome    *string
	Limite  *float64
	Created *time.Time
	Updated *time.Time
}
