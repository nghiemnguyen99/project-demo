package pkg

import (
	"sum/internal/pkg/db/redis"
)

func HanlderRedis(cr, u chan interface{}, key string) {
	for {
		select {
		case data := <-cr:
			CreateDataToRedis(data, key)
		case data := <-u:
			UpdateDataFromRedis(data, key)
		default:
			return
		}
	}
}

func CreateDataToRedis(data interface{}, key string) {
	redis.SetToRedis(key, data)
}

func UpdateDataFromRedis(data interface{}, key string) {
	return
}
