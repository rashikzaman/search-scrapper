package main

import (
	"rashik/search-scrapper/app/http"
	"rashik/search-scrapper/db/migration"
)

func main() {
	migration.Migrate()
	http.InitRouter()
}
