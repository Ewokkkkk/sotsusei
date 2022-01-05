package data

import "time"

type NewUser struct {
	// UserId            in
	PermissionId      string
	UserName          string    `form:"name"`
	UserNameRuby      string    `form:"name_ruby"`
	Password          string    `form:"password"`
	BirthYear         int       `form:"birth_year"`
	BirthMonth        int       `form:"birth_month"`
	BirthDay          int       `form:"birth_day"`
	Email             string    `form:"email"`
	Zipcode           string    `form:"zipcode"`
	PrefectureAddress string    `form:"prefecture"`
	CityAddress       string    `form:"city"`
	StreetAddress     string    `form:"street"`
	PhoneNumber       string    `form:"phonenumber"`
	CreatedDate       time.Time `json:created_date`
	UpdatedDate       time.Time `json:updated_date`
}
