package database

import (
	"context"
	"golang-app/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(conf config.Redis) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Pass,
		DB:       0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
		return nil, err
	}

	log.Println("Connected to Redis Successfully")

	return client, nil
}
