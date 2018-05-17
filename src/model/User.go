package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Pwd      string `json:"-"`
	Phone    string `json:"phone"`
}

type Auth struct {
	Token    string `json:"-"`
	Platform string `json:"platform"`
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}
