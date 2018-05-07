package dao

import (
	"../model"
	"../util"
	"errors"
)

func GetArticle(articleId int64) (model.Article, error) {

	db, err := OpenDB()
	defer db.Close()
	util.CheckErr(err)

	stmt, err := db.Prepare("SELECT * FROM t_article WHERE id=?")
	util.CheckErr(err)

	rows, err := stmt.Query(articleId)
	util.CheckErr(err)
	if rows.Next() {
		var id int64
		var title string
		var content string
		var createdAt int64
		var updatedAt int64
		var uid int64
		rows.Scan(&id, &title, &content, &createdAt, &updatedAt, &uid)
		return model.Article{
			Id:        id,
			Title:     title,
			Content:   content,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Uid:       uid,
		}, nil
	}
	return model.Article{}, errors.New("get articl fail")
}

func DeleteArticle(articleId int64) error {
	db, err := OpenDB()
	defer db.Close()
	util.CheckErr(err)

	stmt, err := db.Prepare("DELETE FROM t_article WHERE id=?")
	util.CheckErr(err)
	_, err = stmt.Exec(articleId)
	return err
}

func GetArticles(page int, limit int) ([]model.Article, error) {
	db, err := OpenDB()
	defer db.Close()
	util.CheckErr(err)

	stmt, err := db.Prepare("SELECT * FROM t_article LIMIT ? OFFSET ?")
	util.CheckErr(err)

	rows, err := stmt.Query(limit, limit*page)
	if err != nil {
		return nil, err
	}
	var list []model.Article
	for rows.Next() {
		var id int64
		var title string
		var content string
		var createdAt int64
		var updatedAt int64
		var uid int64
		rows.Scan(&id, &title, &content, &createdAt, &updatedAt, &uid)
		list = append(list, model.Article{
			Id:        id,
			Title:     title,
			Content:   content,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Uid:       uid,
		})
	}
	return list, nil
}

func PostArticle(uid int64, article *model.Article) (int64, error) {
	db, err := OpenDB()
	defer db.Close()
	util.CheckErr(err)

	stmt, err := db.Prepare("INSERT INTO t_article (title,content,createdAt,updatedAt,uid) VALUES (?,?,?,?,?)")
	util.CheckErr(err)

	res, err := stmt.Exec(article.Title, article.Content, article.CreatedAt, article.UpdatedAt, uid)
	util.CheckErr(err)

	return res.LastInsertId()
}

func PutArticle(articleId int64, article *model.Article) error {
	db, err := OpenDB()
	defer db.Close()
	util.CheckErr(err)

	stmt, err := db.Prepare("UPDATE t_article SET title=?,content=?,createdAt=?,updatedAt=? WHERE id=?")
	util.CheckErr(err)

	_, err = stmt.Exec(article.Title, article.Content, article.CreatedAt, article.UpdatedAt, articleId)
	util.CheckErr(err)
	return err
}
