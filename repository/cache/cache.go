package cache

import (
	"context"
	i "easy_mall/init"
	"time"
)

type Cache struct {
	Ctx context.Context
}

func (c *Cache) Exists(key string) (bool, error) {
	res, err := i.RedisClient.Exists(c.Ctx, key).Result()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	return i.RedisClient.Set(c.Ctx, key, value, expiration).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return i.RedisClient.Get(c.Ctx, key).Result()
}

func (c *Cache) Incr(key string) error {
	return i.RedisClient.Incr(c.Ctx, key).Err()
}

func (c *Cache) LPush(key string, values ...interface{}) error {
	return i.RedisClient.LPush(c.Ctx, key, values).Err()
}

func (c *Cache) BPop(key string, timeout time.Duration) []string {
	return i.RedisClient.BRPop(c.Ctx, timeout, key).Val()
}

func (c *Cache) ZAddIncr(key string, increment float64, member string) error {
	return i.RedisClient.ZIncrBy(c.Ctx, key, increment, member).Err()
}
