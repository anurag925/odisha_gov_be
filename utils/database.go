package utils

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type dbConfig struct {
	URL string
}

var (
	db  *sql.DB
	cfg *dbConfig
)

func InitDB() {
	var err error
	cfg = &dbConfig{URL: "postgres://postgres:@localhost:5432/odisha_dev?sslmode=disable"}
	db, err = sql.Open("pgx", cfg.URL)
	if err != nil {
		Logger().Error("error opening database: ", err)
	}
	err = db.Ping()
	if err != nil {
		Logger().Error("error pinging database: ", err)
	}
	boil.SetDB(db)
}
