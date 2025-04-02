package broker

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type BrokerStore struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

type BrokerMessage struct {
	client *redis.Client
}

func NewBrokerMessage(client *redis.Client) *BrokerMessage {
	return &BrokerMessage{client: client}
}

func (r BrokerMessage) Publish(ctx context.Context, topic string, message BrokerStore) error {

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return r.client.Publish(ctx, topic, body).Err()
}

func (r BrokerMessage) Subscribe(ctx context.Context, topic string, handler func(message BrokerStore)) (err error) {
	pubsub := r.client.Subscribe(ctx, topic)
	channel := pubsub.Channel()

	go func(ch <-chan *redis.Message) {
		msg := <-ch

		var body BrokerStore

		if err = json.Unmarshal([]byte(msg.Payload), &body); err != nil {
			return
		}

		handler(body)

	}(channel)
	return
}
