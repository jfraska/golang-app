package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/jfraska/golang-app/infra/cache"
	"github.com/jfraska/golang-app/infra/database"
	"github.com/jfraska/golang-app/infra/session"
	"github.com/jfraska/golang-app/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	// Db Connect
	db, err := database.ConnectMongoDB(config.Cfg.Database)
	if err != nil {
		panic(err)
	}

	// Storage Connect
	sdb, err := database.ConnectMinio(config.Cfg.Minio)
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

	// Broker initial
	// broker := broker.NewBrokerMessage(rdb)

	// Cache initial
	cache := cache.NewCacheMemory(rdb)

	// Gin Initial;
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	initRoute(router, db, sdb, cache)

	log.Printf("Server started at :%s", config.Cfg.Server.Port)
	if err := router.Run(":" + config.Cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
