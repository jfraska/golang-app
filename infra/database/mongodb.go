package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jfraska/golang-app/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(conf config.Database) (*mongo.Database, error) {

	connectionString := fmt.Sprintf("mongodb://%s:%s", conf.Host, conf.Port)

	clientOptions := options.Client().ApplyURI(connectionString).SetAuth(options.Credential{
		Username: conf.User,
		Password: conf.Pass,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("failed open connection to MongoDB: ", err.Error())
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		client.Disconnect(ctx)
		return nil, err
	}

	log.Println("Connected to MongoDB Successfully")

	return client.Database(conf.Name), nil
}
