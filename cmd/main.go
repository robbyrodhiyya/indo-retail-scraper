package main

import (
	"context"
	"fmt"

	"indo-retail-scraper/internal/db"
	"indo-retail-scraper/internal/repository"
	"indo-retail-scraper/internal/scraper"
)

func main() {
	// connect mongo
	mongoConn, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer mongoConn.Client.Disconnect(context.Background())

	ctx := context.Background()

	// init repository
	categoryRepo := repository.NewCategoryRepository(mongoConn.DB, "alfamart")

	fmt.Println("Scraping Alfamart categories...")

	// scrape categories
	categories, err := scraper.ScrapeAlfamartCategories()
	if err != nil {
		panic(err)
	}

	fmt.Println("Total categories:", len(categories))

	// insert categories
	err = categoryRepo.InsertCategories(ctx, categories)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted categories:", len(categories))
	fmt.Println("Done")
}
