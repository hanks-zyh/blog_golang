package controller

import (
	"../dao"
	"../model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func HandlePostArticle(c *gin.Context) {
	uid := c.PostForm("uid")

	uid64, _ := strconv.ParseInt(uid, 10, 64)
	article := model.Article{
		Uid:     uid64,
		Title:   c.PostForm("title"),
		Content: c.PostForm("content"),
	}

	articleId, e := dao.PostArticle(&article)
	SendResult(c, e, articleId)
}

func HandleDeleteArticle(c *gin.Context) {
	articleId := c.Param("id")
	articleId64, _ := strconv.ParseInt(articleId, 10, 64)
	e := dao.DeleteArticle(articleId64)
	SendResult(c, e, nil)
}

func HandlePutArticle(c *gin.Context) {
	articleId := c.Param("id")
	articleId64, _ := strconv.ParseInt(articleId, 10, 64)
	article, _ := dao.GetArticle(articleId64)
	article.Title = c.PostForm("title")
	article.Content = c.PostForm("content")
	e := dao.PutArticle(&article)
	SendResult(c, e, nil)
}

func HandleGetArticle(c *gin.Context) {
	articleId := c.Param("id")
	articleId64, _ := strconv.ParseInt(articleId, 10, 64)
	article, e := dao.GetArticle(articleId64)
	SendResult(c, e, article)
}

func HandleGetArticles(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	limit := c.DefaultQuery("limit", "20")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	articles, e := dao.GetArticles(pageInt, limitInt)
	SendResult(c, e, articles)
}
