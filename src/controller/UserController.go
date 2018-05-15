package controller

import (
	"../dao"
	"../model"
	"github.com/gin-gonic/gin"
)

func HandleSignup(c *gin.Context) {
	username := c.PostForm("username")
	pwd := c.PostForm("pwd")
	phone := c.PostForm("phone")
	u := model.User{
		Username: username,
		Pwd:      pwd,
		Phone:    phone,
	}
	user, e := dao.SignUp(&u)
	SendResult(c, e, user)
}

func HandleLogin(c *gin.Context) {
	u := model.User{
		Username: c.PostForm("username"),
		Pwd:      c.PostForm("pwd"),
	}
	user, e := dao.Login(&u)
	SendResult(c, e, user)
}
