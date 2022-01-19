package main

import (
	"context"
	"rashik/search-scrapper/app/http"
	"rashik/search-scrapper/app/repository"
	"rashik/search-scrapper/app/scheduler"
	"rashik/search-scrapper/db"
	"rashik/search-scrapper/db/migration"
)

func main() {
	ctx := context.Background()
	migration.Migrate()
	keywordRepository := repository.NewPostgresKeywordRepository(db.GetDb())
	scheduler.ScheduleKeywordParser(ctx, keywordRepository)
	router := http.SetupRouter()
	router.Run()
}
