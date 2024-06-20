package main

import (
	"flag"
	"log"

	"github.com/Lucas-Linhar3s/GerencIA/backend/pkg/jwt"
	pkgLog "github.com/Lucas-Linhar3s/GerencIA/backend/pkg/log"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/adapters/http/server"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/config"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/di"
)

// @title           Api GerencIA backend
// @version         1.0.0
// @description      Api GerencIA backend.
// @termsOfService  http://swagger.io/terms/
// @contact.name   Swagger API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  linhar3s77@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      https://gerencia-v05a.onrender.com
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Load the configuration from file
	config.LoadConfig("config")
	// Get the loaded configuration
	conf := config.GetConfig()

	var envConf = flag.String("conf", "./config/local.json", "config path, eg: -conf ./config/local.json")
	flag.Parse()
	confViper := config.NewConfig(*envConf)

	// Open the database connection using the configuration
	conn, err := database.Open(conf, "server/infrastructure/sqlite", true)
	if err != nil {
		log.Fatal(err)
	}

	jwtJWT := jwt.NewJwt(confViper)

	di.ConfigDi(conn)
	di.UserDi(conn, jwtJWT)
	logger := pkgLog.NewLog(confViper)

	if err := server.InitServer(conf.Server.Port, logger, jwtJWT, confViper); err != nil {
		log.Fatal(err)
	}
}
