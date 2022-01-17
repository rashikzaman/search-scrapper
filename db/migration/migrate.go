package migration

import (
	"rashik/search-scrapper/app/domain"
	"rashik/search-scrapper/db"
)

func Migrate() {
	db.GetDb().AutoMigrate(
		domain.User{},
		domain.Keyword{},
	)
}
