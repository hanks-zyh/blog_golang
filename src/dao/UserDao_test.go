package dao

import (
	"testing"
	"../model"
	"github.com/jinzhu/gorm"
	"fmt"
	"../conf"
	"time"
)

func init() {
	config := conf.AppConfig
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Address,
		config.Mysql.Port,
		config.Mysql.DbName)
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic("failed to connect database")
	}
	model.DB = db
}

func TestLogin(t *testing.T) {
	user := model.User{
		Username: "hanks",
		Pwd:      "123456",
	}
	t.Logf("test login:%v", user)
	up, e := Login(&user)
	if e != nil {
		t.Error(e)
	}
	if up.ID != 0 {
		t.Log("success:" + fmt.Sprint(up.ID))
	} else {
		t.Error("fail to login")
	}
}

func TestSignUp(t *testing.T) {
	user := model.User{
		Username: "hanks" + fmt.Sprint(time.Now().UnixNano()/1000000),
		Pwd:      "123456",
		Phone:    "13586234521",
	}
	t.Logf("test signup:%v", user)
	up, e := SignUp(&user)
	if e != nil {
		t.Error(e)
	}
	if up.ID != 0 {
		t.Log("success:" + fmt.Sprint(up.ID))
	} else {
		t.Fatal("fail to signup")
	}

}
