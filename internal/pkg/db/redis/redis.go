package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func ConnectRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}

	redisClient = client
}

func SetToRedis(key string, val interface{}) error {
	return redisClient.Set(key, val, 0).Err()
}

func GetFromRedis(key string, dest interface{}) error {
	if redisClient.Exists(key).Val() == 0 {
		return nil
	}
	value, err := redisClient.Get(key).Result()
	if err != nil {
		return err
	}
	p := []byte(value)
	return json.Unmarshal(p, dest)

}
