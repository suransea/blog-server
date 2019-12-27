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

type Article struct {
	Id        int64  `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Summary   string `json:"summary" db:"summary"`
	ContentId int64  `json:"content_id" db:"content_id"`
	Ctime     int64  `json:"ctime" db:"ctime"`
	Utime     int64  `json:"utime" db:"utime"`
}

type Content struct {
	Id      int64  `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
	Ctime   int64  `json:"ctime" db:"ctime"`
	Utime   int64  `json:"utime" db:"utime"`
}

type Tag struct {
	Id    int64  `json:"id" db:"id"`
	Tag   string `json:"tag" db:"tag"`
	Ctime int64  `json:"ctime" db:"ctime"`
	Utime int64  `json:"utime" db:"utime"`
}
