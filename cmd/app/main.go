package main

import (
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	log "github.com/sirupsen/logrus"
)

func main() {
	configs.SetUpLogging()
	log.Info("starting app")
}
