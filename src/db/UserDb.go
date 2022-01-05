package db

import (
	"sotsusei/data"
)

//ユーザIdとパスワードをもとにユーザの情報を取得する
func GetUser(email string) data.User {

	d := GormConnect()

	var usr data.User

	d.Find(&usr, "email=?", email)
	return usr
}
