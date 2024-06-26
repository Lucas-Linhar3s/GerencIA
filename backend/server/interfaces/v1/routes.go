package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/interfaces/v1/user"
)

// Router define as rotas da API
func Router(r gin.IRoutes) {
	// User
	r.GET("v1/user/find", user.Find)
	r.POST("v1/user/create", user.Create)
	r.PATCH("v1/user/update", user.Update)
	r.DELETE("v1/user/delete", user.Delete)
}
