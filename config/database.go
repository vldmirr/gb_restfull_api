package config

import (
	"fmt"
	"goland-crud/helper"

	"database/sql"

	_ "github.com/lib/pq" //p
	"github.com/rs/zerolog/log"
)

func DatabaseConnection() *sql.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connect to database")

	return db
}
