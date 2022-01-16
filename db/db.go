package db

import (
	"rashik/search-scrapper/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	if db == nil {
		databaseConfig := config.GetConfig().GetDatabaseConfig()
		dsn := "host=" + databaseConfig.Host + " user=" + databaseConfig.Username +
			" password=" + databaseConfig.Password + " dbname=" + databaseConfig.DatabaseName +
			" port=" + databaseConfig.Port + " sslmode=disable"
		tempDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		} else {
			log.Println("Connected to database")
		}

		db = tempDb
	}
	return db
}
