package dao

import (
	"../model"
	"../util"
	"errors"
	"github.com/gin-gonic/gin/json"
	_ "github.com/go-sql-driver/mysql"
)

func Login(u *model.User) (model.User, error) {
	db, err := OpenDB()
	defer db.Close()
	util.CheckErr(err)

	stmt, err := db.Prepare("SELECT * FROM t_user WHERE username=? AND pwd=?")
	util.CheckErr(err)

	rows, err := stmt.Query(u.Username, u.Pwd)
	util.CheckErr(err)

	if rows.Next() {
		var id int64
		var username string
		var pwd string
		var phone string
		var updatedAt int64
		var createdAt int64
		var authData model.Auth
		rows.Scan(&id, &username, &pwd, &phone, &authData, &updatedAt, &createdAt)
		return model.User{
			Username:  username,
			Phone:     phone,
			Id:        id,
			AuthData:  authData,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}, nil
	}

	return model.User{}, errors.New("sign up fail")
}

func SignUp(u *model.User) (model.User, error) {
	authData, err := json.Marshal(u.AuthData)
	util.CheckErr(err)
	db, err := OpenDB()
	defer db.Close()
	util.CheckErr(err)
	stmt, err := db.Prepare("INSERT INTO t_user (username,pwd,phone,createdAt,updatedAt,authData) VALUES (?,?,?,?,?,?)")
	util.CheckErr(err)
	_, err = stmt.Exec(u.Username, u.Pwd, u.Phone, u.CreatedAt, u.UpdatedAt, string(authData))
	util.CheckErr(err)
	return Login(u)
}
