package config

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Server     Server
	Database   Database
	Encryption Encryption
	Oauth      Oauth
	Redis      Redis
	Minio      Minio
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
	Tz   string
}

type Oauth struct {
	Google OAuthProvider
}

type Redis struct {
	Host string
	Port string
	Pass string
}

type Minio struct {
	Host        string
	Port        string
	MinioAccess string
	MinioSecret string
}

type OAuthProvider struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

type Encryption struct {
	Salt       int64
	JWTSecret  string
	JWTExpires int64
}

var Cfg Config

func Load() {
	fileFlag := flag.String("env", "", "file .env location path absolute")
	flag.Parse()

	var err error
	if *fileFlag != "" {
		err = godotenv.Load(*fileFlag)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("error when load .env: ", err.Error())
	}

	Cfg = Config{
		Server: Server{
			Host: getEnv("SERVER_HOST", "http://localhost"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: Database{
			Host: getEnv("DB_HOST", "127.0.0.1"),
			Port: getEnv("DB_PORT", "3306"),
			User: getEnv("DB_USER", "root"),
			Pass: getEnv("DB_PASS", "mypassword"),
			Name: getEnv("DB_NAME", "golang"),
			Tz:   getEnv("DB_TZ", "Asia/Jakarta"),
		},
		Oauth: Oauth{
			Google: OAuthProvider{
				ClientID:     getEnv("GOOGLE_CLIENT_ID", "oauthsecret"),
				ClientSecret: getEnv("GOOGLE_CLIENT_SECRET", "oauthsecret"),
				CallbackURL:  getEnv("GOOGLE_CALLBACK_URL", "oauthsecret"),
			},
		},
		Redis: Redis{
			Host: getEnv("REDIS_HOST", "127.0.0.1"),
			Port: getEnv("REDIS_PORT", "6379"),
			Pass: getEnv("REDIS_PASSWORD", "redissecret"),
		},
		Minio: Minio{
			Host:        getEnv("MINIO_HOST", "127.0.0.1"),
			Port:        getEnv("MINIO_PORT", "9000"),
			MinioAccess: getEnv("MINIO_ACCESS_KEY", "minioacces"),
			MinioSecret: getEnv("MINIO_SECRET_KEY", "miniosecret"),
		},
		Encryption: Encryption{
			Salt:       getEnvAsInt("SALT", 10),
			JWTSecret:  getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
			JWTExpires: getEnvAsInt("JWT_EXPIRES_IN", 10),
		},
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
