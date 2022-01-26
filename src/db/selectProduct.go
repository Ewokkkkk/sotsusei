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

// トップページ用。最新6件の商品取得
func GetLatestProducts() []data.Product {

	d := GormConnect()

	var products []data.Product
	// var product data.Product

	d.Limit(6).Order("product_id desc").Find(&products)
	// fmt.Printf("%v\n", products)

	return products
}
