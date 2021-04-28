package model

import (
	"database/sql"
	"time"
)

type Tag struct {
	Id    int64  `json:"id" db:"id"`
	Tag   string `json:"tag" db:"tag"`
	Ctime int64  `json:"ctime" db:"ctime"`
	Utime int64  `json:"utime" db:"utime"`
}

func GetTags() (tags []Tag, err error) {
	err = db.Select(&tags, "SELECT id, tag, ctime, utime FROM tag")
	return
}

func GetTag(id int64) (tag Tag, err error) {
	err = db.Get(&tag, "SELECT id, tag, ctime, utime FROM tag WHERE id=?", id)
	if err == sql.ErrNoRows {
		err = ErrNoResults
	}
	return
}

func GetTagArticles(id int64) (articles []Article, err error) {
	err = db.Select(&articles, "SELECT t1.id, t1.title, t1.summary, t1.content_id, t1.ctime, t1.utime FROM "+
		"article AS t1 INNER JOIN article_tag AS t2 ON t1.id = t2.article_id WHERE t2.tag_id=?", id)
	return
}

func AddTag(tag string) (id int64, err error) {
	rst, err := db.Exec("INSERT INTO tag(tag, ctime, utime) VALUES (?, ?, ?)",
		tag, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		return
	}
	return rst.LastInsertId()
}

func UpdateTag(id int64, tag Tag) (err error) {
	_, err = db.Exec("UPDATE tag SET tag=?, utime=? WHERE id=?",
		tag.Tag, time.Now().Unix(), id)
	return
}

func RemoveTag(id int64) (err error) {
	_, err = db.Exec("DELETE FROM tag WHERE id=?", id)
	return
}
