package model

import "time"

type Product struct {
	CreatedAt          time.Time `bson:"created_at"`
	Source             string    `bson:"source"`
	ProductID          string    `bson:"product_id"`
	Name               string    `bson:"name"`
	CategoryID         string    `bson:"category_id"`
	CategoryName       string    `bson:"category_name"`
	SubCategoryID      string    `bson:"subcategory_id"`
	SubCategoryName    string    `bson:"subcategory_name"`
	BrandID            string    `bson:"brand_id"`
	BrandName          string    `bson:"brand_name"`
	Description        string    `bson:"description"`
	ProductInformation []string  `bson:"product_information"`
	Price              float64   `bson:"price"`
	ImageURL           string    `bson:"image_url"`
	URL                string    `bson:"url"`
}
