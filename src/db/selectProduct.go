package db

import (
	"sotsusei/data"
)

func GetProduct(pid string) data.Product {

	d := GormConnect()

	var product data.Product

	d.Find(&product, "product_id=?", pid)
	return product
}

func GetFood(fid int) string {
	d := GormConnect()
	var foodName string
	d.Table("foods").Select("food_name").Where("food_id = ?", fid).Scan(&foodName)

	return foodName
}
