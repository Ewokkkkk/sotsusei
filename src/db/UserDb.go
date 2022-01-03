package db

import (
	"sotsusei/data"
)

//ユーザIdとパスワードをもとにユーザの情報を取得する
func GetUser(email string) data.User {

	db := GormConnect()

	var usr data.User

	db.Find(&usr, "email=?", email)
	return usr
}
