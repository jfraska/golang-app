package session

import (
	"context"
	"encoding/json"
	"golang-app/internal/config"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var Store *sessionStore

type Session struct {
	Name   string    `json:"name"`
	UserID uuid.UUID `json:"user_id"`
}

type sessionStore struct {
	client *redis.Client
}

func NewSession(client *redis.Client) *sessionStore {
	return &sessionStore{client: client}
}

func (r *sessionStore) Set(ctx context.Context, id string, session Session) error {
	body, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, id, body, time.Duration(config.Cfg.Encryption.JWTExpires)*time.Minute).Err()
}

func (r *sessionStore) Get(ctx context.Context, id string) (Session, error) {
	var session Session

	body, err := r.client.Get(ctx, id).Bytes()
	if err != nil {
		return session, err
	}

	if err := json.Unmarshal(body, &session); err != nil {
		return session, err
	}

	return session, nil
}

func (r *sessionStore) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
