# indo-retail-scraper

Simple Golang scraper for Indonesian retail and grocery stores.

## Features
- **Multi-channel**: Scrapes product listings and pricing from various Indonesian retail platforms.
- **Direct Storage**: Saves scraped data directly to a **MongoDB Atlas** cluster.
- **Concurrent**: Built with Go for high-performance and efficient processing.
- **Modular**: Designed to easily add new retail channels in the future.

## Prerequisites
- Go 1.26 or later.
- A MongoDB Atlas account and an active cluster.

## Setup
1. Clone this repository:
   ```bash
   git clone https://github.com/robbyrodhiyya/indo-retail-scraper.git
   cd indo-retail-scraper
   ```

2. Create a `.env` file in the root directory and add your credentials:
   ```env
   MONGO_URI=your_mongodb_connection_string
   DB_NAME=your_db_name_string
   ```

## Usage
Run the scraper using the following command:
```bash
go run ./cmd [source]
```

*Or using Docker:*
```bash
docker compose run --rm app go run ./cmd [source]
```

Replace `[source]` with one of the currently supported channels:
- `alfagift`
- `indomaret`

## Data Schema
The collected data typically includes:
- **Product Name**: Name of the item.
- **Current Price**: The latest price after discount.
- **Original Price**: Base price before discount (if available).
- **Category**: Product category or tag.
- **Product URL**: Link to the product page.
- **Image URL**: Direct link to the product image.
- **Store Name**: The source marketplace name.
- **Scraped At**: Timestamp of the scraping process.
