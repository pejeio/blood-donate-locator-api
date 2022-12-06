package configs

import (
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	log "github.com/sirupsen/logrus"
)

func AutoMigrate() {
	// configs.Db().Migrator().DropTable("locations")
	Db().AutoMigrate(&types.Location{})
	log.Println("👍 Database migration complete")
}
