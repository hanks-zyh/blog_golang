package controller

import (
	"../dao"
	"../model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"fmt"
)

func HandlePostArticle(c *gin.Context) {
	uid := c.PostForm("uid")
	article := model.Article{
		Title:     c.PostForm("title"),
		Content:   c.PostForm("content"),
		CreatedAt: time.Now().UnixNano() / 1000000,
		UpdatedAt: time.Now().UnixNano() / 1000000,
	}

	uid64, _ := strconv.ParseInt(uid, 10, 64)
	articleId, e := dao.PostArticle(uid64, &article)
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
	article.UpdatedAt = time.Now().UnixNano() / 1000000
	e := dao.PutArticle(articleId64, &article)
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
