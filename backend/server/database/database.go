package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Lucas-Linhar3s/GerencIA/backend/server/config"
)

type Database struct {
	db                 *sql.DB
	transactionTimeout int
	Builder            sq.StatementBuilderType
}

func Open(c *config.Config, dir string, driver string) (database *Database, err error) {
	var db *sql.DB

	if driver == "mysql" || driver == "postgres" {
		driverConfig := stdlib.DriverConfig{
			ConnConfig: pgx.ConnConfig{
				RuntimeParams: map[string]string{
					//Verificar
					"application_name": "GerencIA",
					"DateStyle":        "ISO",
					"IntervalStyle":    "iso_8601",
					// TODO:
					"search_path": "public",
				},
			},
		}
		stdlib.RegisterDriverConfig(&driverConfig)

		db, err = sql.Open("pgx", driverConfig.ConnectionString(
			*c.Databases.Nick+
				"://"+
				*c.Databases.Username+
				":"+
				*c.Databases.Password+
				"@"+
				*c.Databases.Host+
				":"+
				*c.Databases.Port+
				"/"+
				*c.Databases.Name))
		if err != nil {
			return nil, err
		}
	} else if driver == "sqlite" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path := filepath.Join(currentDir, fmt.Sprintf("%s/%s", dir, "/databaseApp.db"))

		db, err = sql.Open("sqlite3", path)
		if err != nil {
			return nil, err
		}
	} else {
		panic("unknown db driver")
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(*c.Databases.MaxIdle)
	db.SetMaxOpenConns(*c.Databases.MaxConn)
	db.SetConnMaxLifetime(time.Second * 60)

	return &Database{
		db:                 db,
		transactionTimeout: *c.Databases.TransactionTimeout,
		Builder:            sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db),
	}, nil
}

func (p *Database) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
