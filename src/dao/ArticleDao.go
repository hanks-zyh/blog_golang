package dao

import (
	"../model"
	"errors"
)

func GetArticle(articleId int64) (model.Article, error) {
	db := model.DB
	var article model.Article
	if e := db.First(&article, articleId).Error; e != nil {
		return model.Article{}, errors.New("article does't exist")
	}
	return article, nil
}

func DeleteArticle(articleId int64) error {
	db := model.DB
	if e := db.Delete(&model.Article{}, "id=?", articleId).Error; e != nil {
		return errors.New("delete article fail")
	}
	return nil
}

func GetArticles(page int, limit int) ([]model.Article, error) {
	db := model.DB

	var list []model.Article
	db.Limit(limit).Offset(page * limit).Find(&list)
	return list, nil
}

func PostArticle(article *model.Article) (uint, error) {
	db := model.DB
	if e := db.First(&model.User{}, article.Uid).Error; e != nil {
		return 0, errors.New("user is not exist")
	}

	if err := db.Create(article).Error; err != nil {
		return 0, err
	}
	return article.ID, nil
}

func PutArticle(article *model.Article) error {
	db := model.DB

	if e := db.Save(article).Error; e != nil {
		return errors.New("update article error")
	}
	return nil
}
