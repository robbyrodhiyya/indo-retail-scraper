package scraper

import (
	"encoding/json"
	"net/http"
	"time"

	"indo-retail-scraper/internal/model"
	"indo-retail-scraper/internal/utils"
)

type AlfamartCategoryResponse struct {
	Categories []struct {
		CategoryID   string `json:"categoryId"`
		CategoryName string `json:"categoryName"`

		SubCategories []struct {
			CategoryID   string `json:"categoryId"`
			CategoryName string `json:"categoryName"`
		} `json:"subCategories"`
	} `json:"categories"`
}

func ScrapeAlfamartCategories() ([]model.Category, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(
		"GET",
		"https://webcommerce-gw.alfagift.id/v2/categories",
		nil,
	)
	if err != nil {
		return nil, err
	}

	utils.SetBrowserHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response AlfamartCategoryResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	var categories []model.Category

	for _, category := range response.Categories {

		// jika tidak punya subcategory
		if len(category.SubCategories) == 0 {
			categories = append(categories, model.Category{
				Source:          "alfamart",
				CategoryID:      category.CategoryID,
				CategoryName:    category.CategoryName,
				SubCategoryID:   "",
				SubCategoryName: "",
			})

			continue
		}

		// flatten subcategories
		for _, sub := range category.SubCategories {
			categories = append(categories, model.Category{
				Source:          "alfamart",
				CategoryID:      category.CategoryID,
				CategoryName:    category.CategoryName,
				SubCategoryID:   sub.CategoryID,
				SubCategoryName: sub.CategoryName,
			})
		}
	}

	return categories, nil
}
