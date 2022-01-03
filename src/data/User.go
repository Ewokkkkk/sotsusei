package data

import "time"

// type User struct {
// 	gorm.Model
// 	UserId     string `gorm:"primary_key"`
// 	Password   string
// 	RegistDate string
// 	LastLogin  string
// }

type User struct {
	UserId            int       `json:user_id`
	PermissionId      string    `json:permission_id`
	FamilyName        string    `json:family_name`
	FirstName         string    `json:first_name`
	FamilyNameRuby    string    `json:family_name_ruby`
	FirstNameRuby     string    `json:first_name_ruby`
	Password          string    `json:password`
	BirthYear         int       `json:birth_year`
	BirthMonth        int       `json:birth_month`
	BirthDay          int       `json:birth_day`
	Email             string    `json:email`
	Zipcode           string    `json:zipcode`
	PrefectureAddress string    `json:prefecture_address`
	StreetAddress     string    `json:street_address`
	PhoneNumber       string    `json:phone_number`
	CreatedDate       time.Time `json:created_date`
	UpdatedDate       time.Time `json:updated_date`
}
