package cache

import (
	"context"
	"easy_mall/init"
	"time"
)

type Cache struct {
	Ctx context.Context
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	return init.RedisClient.Set(c.Ctx, key, value, expiration).Err()
}

func (c *Cache) Get(key string) string {
	return init.RedisClient.Get(c.Ctx, key).Val()
}

func (c *Cache) Incr(key string) error {
	return init.RedisClient.Incr(c.Ctx, key).Err()
}

func (c *Cache) LPush(key string, values ...interface{}) error {
	return init.RedisClient.LPush(c.Ctx, key, values).Err()
}

func (c *Cache) BPop(key string, timeout time.Duration) []string {
	return init.RedisClient.BRPop(c.Ctx, timeout, key).Val()
}

func (c *Cache) ZAddIncr(key string, increment float64, member string) error {
	return init.RedisClient.ZIncrBy(c.Ctx, key, increment, member).Err()
}
