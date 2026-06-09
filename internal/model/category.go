package model

type Category struct {
	Source          string `bson:"source"`
	CategoryID      string `bson:"category_id"`
	CategoryName    string `bson:"category_name"`
	SubCategoryID   string `bson:"subcategory_id"`
	SubCategoryName string `bson:"subcategory_name"`
}
