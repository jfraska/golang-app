package main

import (
	"golang-app/infra/database"
	"golang-app/infra/session"
	"golang-app/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	// Db Connect
	db, err := database.ConnectMongoDB(config.Cfg.Database)
	if err != nil {
		panic(err)
	}

	// Redis Connect
	rdb, err := database.ConnectRedis(config.Cfg.Redis)
	if err != nil {
		panic(err)
	}

	// Session initial
	session.Store = session.NewSession(rdb)

	// Gin Initial;
	router := gin.Default()

	initRoute(router, db)

	log.Printf("Server started at :%s", config.Cfg.Server.Port)
	if err := router.Run(":" + config.Cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
