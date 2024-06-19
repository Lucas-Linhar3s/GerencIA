package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/service"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/di"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/docs"
)

// InitServer inicializa o servidor Gin
func InitServer(port string) (err error) {
	// Cria o servidor Gin
	r := gin.Default()

	r.Use(
		gin.Logger(),
	)

	// Definir suas rotas aqui
	cardsRoutes(r, di.GetServices())
	internalRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found " + " : " + c.Request.URL.String()})
	})

	// Inicia o servidor
	if err = r.Run(port); err != nil {
		return err
	}

	return
}

func cardsRoutes(r *gin.Engine, handlers service.ICardServices) {
	course := r.Group("card")
	{
		course.POST("/create", handlers.Create)
		course.GET("/list", handlers.Find)
		course.PATCH("/update", handlers.Update)
		course.DELETE("/delete/:id", handlers.Delete)
	}
}

func internalRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	api := r.Group("api")

	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
