package scraper

import "indo-retail-scraper/internal/model"

type Scraper interface {
	Scrape() ([]model.Product, error)
	GetName() string
}
