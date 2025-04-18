package session

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jfraska/golang-app/internal/config"

	"github.com/redis/go-redis/v9"
)

var Store *session

type SessionStore struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image,omitempty"`
}

type session struct {
	client *redis.Client
}

func NewSession(client *redis.Client) *session {
	return &session{client: client}
}

func (r *session) Set(ctx context.Context, ID string, session SessionStore) error {
	body, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, ID, body, time.Duration(config.Cfg.Encryption.JWTExpires)*time.Hour).Err()
}

func (r *session) Get(ctx context.Context, ID string) (SessionStore, error) {
	var session SessionStore

	body, err := r.client.Get(ctx, ID).Bytes()
	if err != nil {
		return session, err
	}

	if err := json.Unmarshal(body, &session); err != nil {
		return session, err
	}

	return session, nil
}

func (r *session) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
