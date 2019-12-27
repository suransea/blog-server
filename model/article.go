package model

import "database/sql"

func GetArticles() (articles []Article, err error) {
	err = db.Select(&articles, "SELECT id, title, summary, content_id, ctime, utime FROM article")
	return
}

func GetArticle(id int64) (article Article, err error) {
	err = db.Get(&article, "SELECT id, title, summary, content_id, ctime, utime FROM article WHERE id=?", id)
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
