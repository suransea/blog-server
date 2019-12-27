package model

import (
	"time"
)

func GetTags() (tags []Tag, err error) {
	err = db.Select(&tags, "SELECT id, tag, ctime, utime FROM tag")
	return
}

func GetTag(id int64) (tag Tag, err error) {
	err = db.Get(&tag, "SELECT id, tag, ctime, utime FROM tag WHERE id=?", id)
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
