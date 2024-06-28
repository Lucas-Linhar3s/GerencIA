package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/interfaces/v1/card"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/interfaces/v1/debt"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/interfaces/v1/user"
)

// Router define as rotas da API
func Router(r gin.IRoutes) {
	// User
	r.GET("v1/user/find", user.Find)
	r.POST("v1/user/create", user.Create)
	r.PATCH("v1/user/update", user.Update)
	r.DELETE("v1/user/delete", user.Delete)
	//Card
	r.GET("v1/card/find", card.Find)
	r.POST("v1/card/create", card.Create)
	r.PATCH("v1/card/update", card.Update)
	r.DELETE("v1/card/delete", card.Delete)
	// Debt
	r.GET("v1/debt/find", debt.Find)
	r.POST("v1/debt/create", debt.Create)
	r.PATCH("v1/debt/update", debt.Update)
	r.DELETE("v1/debt/delete", debt.Delete)
}
