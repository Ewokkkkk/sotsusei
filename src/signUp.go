package main

import (
	"sotsusei/data"
	"sotsusei/db"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func signUp(user data.NewUser) error {
	log.Printf("%+v\n", user)

	passwordEncrypt, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	log.Print(string(passwordEncrypt))
	user.Password = string(passwordEncrypt)
	user.UserId = 1
	user.PermissionId = "1"
	d := db.GormConnect()
	// Insert処理
	if err := d.Table("users").Create(user).Error; err != nil {
		return err
	}
	return nil
}
