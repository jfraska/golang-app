package database

import (
	"fmt"
	"log"

	"github.com/jfraska/golang-app/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectMinio(conf config.Minio) (*minio.Client, error) {

	connectionString := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

	client, err := minio.New(connectionString, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.MinioAccess, conf.MinioSecret, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatal("failed open connection to Minio: ", err.Error())
		return nil, err
	}

	log.Println("Connected to Minio Successfully")

	return client, nil
}
