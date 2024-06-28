package card

import "time"

// CardModel is the model for card
type CardModel struct {
	ID      *string
	Nome    *string
	Limite  *float64
	Created *time.Time
	Updated *time.Time
}
