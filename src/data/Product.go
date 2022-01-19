package data

import "time"

type Product struct {
	ProductId       *int   `json:product_id`
	FoodId          int    `json:food_id`
	SellerId        int    `json:seller_id`
	ProductName     string `json:product_name`
	ProductLocality string `json:product_locality`
	ProductImg1     string `json:product_img1`
	ProductImg2     string `json:product_img2`
	ProductImg3     string `json:product_img3`
	// ProductImg         []string  `json:product_img, product_img2, product_img3`
	ProductPrice       string    `json:product_price`
	ProductStock       string    `json:product_stock`
	ProductDescription string    `json:product_description`
	Deadline           time.Time `json:deadline`
	DeadlineStr        string
	CreatedDate        time.Time `json:created_date`
	UpdatedDate        time.Time `json:updated_date`
}
