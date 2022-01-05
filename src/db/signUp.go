package db

import (
	"sotsusei/data"
	"time"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func IsDuplicate(user data.NewUser) bool {
	log.Printf("%+v\n", user)
	d := GormConnect()
	// 重複する行があればtrue
	if result := d.Table("users").Find(&user, "email=?", user.Email); result.RowsAffected > 0 {
		return true
	}
	return false
}
func SignUp(user data.NewUser) error {
	log.Printf("%+v\n", user)

	passwordEncrypt, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// log.Print(string(passwordEncrypt))
	user.Password = string(passwordEncrypt)
	user.PermissionId = "1"
	user.CreatedDate = time.Now()
	user.UpdatedDate = time.Now()
	d := GormConnect()
	// Insert処理
	if err := d.Table("users").Create(user).Error; err != nil {
		return err
	}
	return nil
}
