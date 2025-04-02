package database

import (
	"context"
	"log"
	"time"

	"github.com/jfraska/golang-app/internal/config"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(conf config.Redis) (*redis.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Pass,
		DB:       0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
		return nil, err
	}

	log.Println("Connected to Redis Successfully")

	return client, nil
}
