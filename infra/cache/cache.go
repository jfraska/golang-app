package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type CacheStore struct {
	Content json.RawMessage `json:"value"`
}

type CacheMemory struct {
	client *redis.Client
}

func NewCacheMemory(client *redis.Client) *CacheMemory {
	return &CacheMemory{client: client}
}

func (r CacheMemory) Set(ctx context.Context, key string, value CacheStore) (err error) {
	body, err := json.Marshal(value)
	if err != nil {
		return
	}

	return r.client.Set(ctx, key, body, 0).Err()
}

func (r CacheMemory) Get(ctx context.Context, key string) (value CacheStore, err error) {
	body, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &value); err != nil {
		return
	}

	return
}

func (r CacheMemory) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
