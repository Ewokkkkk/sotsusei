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
