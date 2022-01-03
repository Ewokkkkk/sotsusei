package data

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	food_name string
}
