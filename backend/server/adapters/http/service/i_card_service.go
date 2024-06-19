package service

import "github.com/gin-gonic/gin"

type ICardServices interface {
	Create(ctx *gin.Context)
	Find(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}
