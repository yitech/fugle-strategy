package pubsub

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Publisher struct {
	client *redis.Client
}

func NewPublisher(client *redis.Client) *Publisher {
	return &Publisher{
		client: client,
	}
}

// Publish sends data of any Go type to the specified topic
func (p *Publisher) Publish(ctx context.Context, topic string, message interface{}) error {
	err := p.client.Publish(ctx, topic, message).Err()
	if err != nil {
		log.Printf("Failed to publish message to topic %s: %v", topic, err)
		return err
	}
	log.Printf("Published message to topic %s: %v", topic, message)
	return nil
}
