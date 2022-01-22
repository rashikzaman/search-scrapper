package logger

import (
	"rashik/search-scrapper/config"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = nil

func GetLog() *logrus.Logger {
	if log == nil {
		log = logrus.New()
		if config.GetConfig().GetAppEnv() == "development" {
			log.SetLevel(logrus.DebugLevel)
		} else {
			log.SetLevel(logrus.InfoLevel)
		}
	}
	return log
}
