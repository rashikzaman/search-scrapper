package main

import (
	"context"
	"rashik/search-scrapper/app/http"
	"rashik/search-scrapper/app/repository"
	"rashik/search-scrapper/app/scheduler"
	"rashik/search-scrapper/config"
	"rashik/search-scrapper/db"
	"rashik/search-scrapper/db/migration"
)

func main() {
	migration.Migrate()
	scheduler.ScheduleKeywordParser(context.Background(), repository.NewPostgresKeywordRepository(db.GetDb()))
	router := http.SetupRouter()
	router.Run(":" + config.GetConfig().GetServerPort())
}
