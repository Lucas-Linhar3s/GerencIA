package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Lucas-Linhar3s/GerencIA/backend/pkg/jwt"
	"github.com/Lucas-Linhar3s/GerencIA/backend/pkg/log"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/service"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/di"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/docs"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/middleware"
)

// InitServer inicializa o servidor Gin
func InitServer(port string, logger *log.Logger,
	jwt *jwt.JWT, conf *viper.Viper) (err error) {
	// Cria o servidor Gin
	r := gin.Default()

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		middleware.SignMiddleware(logger, conf),
	)

	// Definir suas rotas com autenticação aqui
	{
		cardsRoutes(r, di.GetServices(), logger, jwt)
	}

	// Configura o Swagger e rotas publicas
	{
		docs.SwaggerInfo.BasePath = "/"
		api := r.Group("api")

		{
			api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
			api.POST("/v1/login", di.GetUserService().Login)
			api.POST("/v1/register", di.GetUserService().Register)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found " + " : " + c.Request.URL.String()})
	})

	// Inicia o servidor
	if err = r.Run(port); err != nil {
		return err
	}

	return
}

func cardsRoutes(r *gin.Engine, handlers service.ICardServices, logger *log.Logger,
	jwt *jwt.JWT) {
	course := r.Group("card").Use(
		middleware.StrictAuth(jwt, logger),
	)
	{
		course.POST("/create", handlers.Create)
		course.GET("/list", handlers.Find)
		course.PATCH("/update", handlers.Update)
		course.DELETE("/delete/:id", handlers.Delete)
	}
}
