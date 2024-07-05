package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/config"
	v1 "github.com/Lucas-Linhar3s/GerencIA/backend/server/interfaces/v1"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/middleware"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/pkg/jwt"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/pkg/log"
)

// InitServer inicializa o servidor Gin
func InitServer(port string, logger *log.Logger,
	jwt *jwt.JWT, conf *config.Config) (err error) {
	// Cria o servidor Gin
	r := gin.Default()

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)

	// Definir suas rotas com autenticação aqui
	api := r.Group("api").Use(
	// middleware.StrictAuth(jwt, logger),
		middleware.SignMiddleware(logger, conf),
	)
	{
		v1.Router(api)
	}

	// Configura o Swagger e rotas publicas
	public := r.Group("public")
	{
		// docs.SwaggerInfo.BasePath = "/"
		{
			public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
