package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title     string `json:"title"`
	Content   string `json:"content"`
	Uid       int64  `json:"uid"`
}
