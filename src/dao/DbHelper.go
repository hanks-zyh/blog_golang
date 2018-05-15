package dao

import (
	"../model"
	"github.com/jinzhu/gorm"
)

func init() {
	// Migrate the schema
	db, _ := OpenDB()
	db.AutoMigrate(&model.User{},&model.Article{})
	defer db.Close()
}

func OpenDB() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog_go?parseTime=true")
}
