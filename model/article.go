package model

import "database/sql"

type Article struct {
	Id      int64  `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Summary string `json:"summary" db:"summary"`
	Ctime   int64  `json:"ctime" db:"ctime"`
	Utime   int64  `json:"utime" db:"utime"`
}

type Content struct {
	Id      int64  `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
	Ctime   int64  `json:"ctime" db:"ctime"`
	Utime   int64  `json:"utime" db:"utime"`
}

func GetArticles() (articles []Article, err error) {
	err = db.Select(&articles, "SELECT id, title, summary, ctime, utime FROM article")
	return
}

func GetArticle(id int64) (article Article, err error) {
	err = db.Get(&article, "SELECT id, title, summary, ctime, utime FROM article WHERE id=?", id)
	if err == sql.ErrNoRows {
		err = ErrNoResults
	}
	return
}

func GetArticleContent(id int64) (content Content, err error) {
	row := db.QueryRow("SELECT content_id FROM article where id=?", id)
	var cid int64
	err = row.Scan(&cid)
	if err == sql.ErrNoRows {
		err = ErrNoResults
		return
	}
	err = db.Get(&content, "SELECT id, content, ctime, utime FROM content WHERE id=?", cid)
	if err == sql.ErrNoRows {
		err = ErrNoResults
	}
	return
}

func GetArticleTags(id int64) (tags []Tag, err error) {
	err = db.Select(&tags, "SELECT t1.id, t1.tag, t1.ctime, t1.utime FROM "+
		"tag AS t1 INNER JOIN article_tag AS t2 ON t1.id = t2.tag_id WHERE t2.article_id=?", id)
	return
}
