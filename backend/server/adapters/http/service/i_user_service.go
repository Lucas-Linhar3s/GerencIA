package service

import "github.com/gin-gonic/gin"

type IUserService interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}
