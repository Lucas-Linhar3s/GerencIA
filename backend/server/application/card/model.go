package card

import "time"

// CardReq struct for request
type CardReq struct {
	Id     *string  `json:"id"`
	Nome   *string  `json:"nome" binding:"required" example:"John Doe"`
	Limite *float64 `json:"limite" binding:"required" example:"100"`
}

// CardRes struct for response
type CardRes struct {
	ID      *string  `json:"id"`
	Nome    *string  `json:"nome"`
	Limite  *float64 `json:"limite"`
	Created *time.Time
	Updated *time.Time
}

type ParamsCard struct {
	Id   *string `form:"id"`
	Nome *string `form:"nome"`
}
