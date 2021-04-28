package model

import (
	"blog-server/config"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Global variable `db` for database operations.
var db *sqlx.DB

// `ErrNoResults` is returned when `Get` return a `ErrNoRows`.
var ErrNoResults = errors.New("model: no results found")

// Initialize the database handle with the config.
func Init(cfg config.Database) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&loc=%s",
		cfg.Username, cfg.Password, cfg.Protocol, cfg.Host, cfg.Port, cfg.Dbname, cfg.Charset, cfg.Loc)

	var err error
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic("Error occurred in opening database.")
	}
}
