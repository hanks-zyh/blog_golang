package dao

import (
	"../model"
	"../util"
	_ "github.com/go-sql-driver/mysql"
	"errors"
)

func Login(u *model.User) (model.User, error) {
	db, err := OpenDB()
	defer db.Close()
	util.CheckErr(err)
	var user model.User
	if e := db.Where("username=? AND pwd=?", u.Username, u.Pwd).First(&user).Error; e!=nil{
		return model.User{},errors.New("username or password is wrong")
	}
	return user,nil
}

func SignUp(u *model.User) (model.User, error) {

	db, err := OpenDB()
	util.CheckErr(err)
	defer db.Close()

	// Read
	var user model.User
	if e := db.Where("username=?", u.Username).First(&user).Error; e == nil{
		return model.User{}, errors.New("username has exist")
	}

	db.Create(u)
	return Login(u)
}
