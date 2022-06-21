package pkg

import (
	"sum/internal/pkg/db/redis"
)

func HanlderRedis(cr, u chan interface{}, key string) error {
	for {
		select {
		case data := <-cr:
			CreateDataToRedis(data, key)
		case data := <-u:
			UpdateDataFromRedis(data, key)
		default:
			return nil
		}
	}
}

func CreateDataToRedis(data interface{}, key string) error {
	return redis.SetToRedis(key, data)
}

func UpdateDataFromRedis(data interface{}, key string) error {
	return nil
}
