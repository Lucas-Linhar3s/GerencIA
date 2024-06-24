package main

import (
	"flag"
	"log"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/config"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/database"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/di"
	pgkHttp "github.com/Lucas-Linhar3s/GerencIA/backend/server/pkg/http"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/pkg/jwt"
	pkgLog "github.com/Lucas-Linhar3s/GerencIA/backend/server/pkg/log"
	"github.com/Lucas-Linhar3s/GerencIA/backend/server/pkg/viper"
)

func main() {
	// Load config file
	config.LoadConfig("../../config")
	// Get config struct
	conf := config.GetConfig()

	var envConf = flag.String("conf", "../../config/local.api.json", "config path, eg: -conf ../../config/local.api.json")
	flag.Parse()
	confViper := viper.NewViper(*envConf)

	// Open the database connection using the configuration
	conn, err := database.Open(conf, "server/infrastructure/sqlite", "sqlite")
	if err != nil {
		log.Fatal(err)
	}

	jwtJWT := jwt.NewJwt(conf)

	di.ConfigDI(conn, jwtJWT)
	logger := pkgLog.NewLog(confViper)

	if err := pgkHttp.InitServer(*conf.Server.Port, logger, jwtJWT, conf); err != nil {
		log.Fatal(err)
	}
}
