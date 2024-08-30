package noSql

import (
	"context"
	_ "fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

// RedisClient redis connection
var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if RedisClient == nil {
		panic("noSql 链接失败")
	}
}

func SetString(key string, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

func RemoveString(key string) error {
	return RedisClient.Del(ctx, key).Err()
}
