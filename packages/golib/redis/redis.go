package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct{
	client *redis.Client
}

func NewConnection() *redis.Client{
	client:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

func (r *RedisClient)PublishToRedis(ctx context.Context,channel string,data []byte) *redis.IntCmd{
	publish:=r.client.Publish(ctx,channel,data)
	return publish
}

func (r *RedisClient)SubcribeToRedis(ctx context.Context,channel string)  *redis.PubSub{
	subscribe:=r.client.Subscribe(ctx,channel);
	return subscribe
}

func (r *RedisClient) AddSymbol(ctx context.Context,channel string,symbol string)*redis.IntCmd {
	add:=r.client.SAdd(ctx,channel,symbol)
	return add
}

func (r *RedisClient) RemoveSymbol(ctx context.Context,channel string,symbol string)*redis.IntCmd{
	remove:=r.client.SRem(ctx,channel,symbol)
	return remove
}

