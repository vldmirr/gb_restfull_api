package config

import (
	"fmt"
	"goland-crud/helper"

	"database/sql"

	_ "github.com/lib/pq" //pg_driver
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "123456"
	dbName   = "postgres"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connect to database")

	return db
}
