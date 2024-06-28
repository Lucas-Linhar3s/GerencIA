package user

import "time"

type UserReq struct {
	ID       string `json:"id"`
	Nome     string `json:"nome" binding:"required" example:"John Doe"`
	Whatsapp string `json:"whatsapp" binding:"required" example:"21999999999"`
}

type UserRes struct {
	ID        *string    `json:"id"`
	Nome      *string    `json:"nome"`
	Whatsapp  *string    `json:"whatsapp"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Params struct {
	Nome     *string `form:"nome"`
	Whatsapp *string `form:"whatsapp"`
	Id       *string `form:"id"`
}
