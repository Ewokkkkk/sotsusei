package data

import "time"

type NewUser struct {
	UserId            int
	PermissionId      string
	FamilyName        string `form:"name"`
	FirstName         string
	FamilyNameRuby    string `form:"name_ruby"`
	FirstNameRuby     string
	Password          string    `form:"password"`
	BirthYear         int       `form:"birth_year"`
	BirthMonth        int       `form:"birth_month"`
	BirthDay          int       `form:"birth_day"`
	Email             string    `form:"email"`
	Zipcode           string    `form:"zipcode"`
	PrefectureAddress string    `form:"prefecture"`
	StreetAddress     string    `form:"street"`
	PhoneNumber       string    `form:"phonenumber"`
	CreatedDate       time.Time `json:created_date`
	UpdatedDate       time.Time `json:updated_date`
}
