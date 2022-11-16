package configs

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Brussels", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	log.Println("🚀 Connected successfully to the database")
}

func Db() *gorm.DB {
	return dbInstance
}
