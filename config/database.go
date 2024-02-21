package config

import (
	"database/sql"
	"fmt"
	"github.com/abbasfisal/go-api/helper"
	_ "github.com/lib/pq" //posgres driver
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "password"
	dbName   = "test"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbName)
	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	log.Info().Msg("Database Connected !")

	return db
}
