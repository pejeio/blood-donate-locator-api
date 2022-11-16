package models

import (
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	log "github.com/sirupsen/logrus"
)

func AutoMigrate() {
	configs.Db().AutoMigrate(&Location{})
	log.Println("👍 Database migration complete")
}
