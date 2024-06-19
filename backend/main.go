package main

import (
	"log"

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

	// Open the database connection using the configuration
	conn, err := database.Open(conf, "server/adapters/sqlite", true)
	if err != nil {
		log.Fatal(err)
	}

	di.ConfigDi(conn)

	if err := server.InitServer(conf.Server.Port); err != nil {
		log.Fatal(err)
	}
}
