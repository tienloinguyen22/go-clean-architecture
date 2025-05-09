package event

import (
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/event"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type PubSub struct {
	client *redis.Client
	ctx    context.Context
	wg     sync.WaitGroup
}

func NewPubSub(config *RedisConfig) (*PubSub, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &PubSub{
		client: client,
		ctx:    context.Background(),
	}, nil
}

func (p *PubSub) Publish(channel string, event event.Event) error {
	return p.client.Publish(p.ctx, channel, event).Err()
}

func (p *PubSub) Close() {
	p.wg.Wait()
	p.client.Close()
}
