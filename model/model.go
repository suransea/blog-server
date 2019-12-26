package model

import (
	"blog-server/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

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
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Tag struct {
	Id    int64  `json:"id" db:"id"`
	Tag   string `json:"tag" db:"tag"`
	Ctime int64  `json:"ctime" db:"ctime"`
	Utime int64  `json:"utime" db:"utime"`
}
