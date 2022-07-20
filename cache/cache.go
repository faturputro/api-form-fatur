package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host     string
	port     int
	db       int
	password string
	ex       time.Duration
}

type AppCache interface {
	Set(key string, data interface{}, duration time.Duration) error
	Get(key string, returnType interface{}) (interface{}, error)
	Del(key string) error
	DelPattern(keyPattern string) error
	Exist(key string) bool
	Pipe() redis.Pipeliner
}

var ctx = context.Background()

func NewRedisCache(host string, port int, db int, pass string) AppCache {
	return &redisCache{
		host:     host,
		port:     port,
		password: pass,
		db:       db,
	}
}

func (c *redisCache) getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.host, c.port),
		Password: c.password,
		DB:       c.db,
	})

	log.Println("Redis initialized.")

	return client
}

func (c *redisCache) Get(key string, returnType interface{}) (interface{}, error) {
	client := c.getClient()

	val, err := client.Get(ctx, key).Result()

	if err != nil {
		if returnType == nil {
			err = errors.New("Cache key doesn't exist: " + key)
			return nil, err
		}
		err = errors.New("Cache key doesn't exist: " + key)
		v := reflect.ValueOf(returnType)
		return reflect.Zero(v.Type()).Interface(), err
	}

	json.Unmarshal([]byte(val), &returnType)

	return val, err
}

func (c *redisCache) Set(key string, data interface{}, ex time.Duration) error {
	client := c.getClient()

	json, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	err = client.Set(ctx, key, json, ex*time.Second).Err()

	if err != nil {
		return err
	}
	return nil
}

func (c *redisCache) Del(key string) error {
	client := c.getClient()

	_, err := client.Del(ctx, key).Result()
	return err
}

func (c *redisCache) DelPattern(key string) error {
	return nil
}

func (c *redisCache) Exist(key string) bool {
	client := c.getClient()

	_, err := client.Get(ctx, key).Result()
	return err != nil
}

func (c *redisCache) Pipe() redis.Pipeliner {
	pipe := c.getClient().Pipeline()

	return pipe
}
