package database

import (
	"fmt"
	"log"

	"github.com/jfraska/golang-app/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(conf config.Database) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Host, conf.User, conf.Pass, conf.Name, conf.Port, conf.Tz)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed open connection to db: ", err.Error())
	}

	return
}
