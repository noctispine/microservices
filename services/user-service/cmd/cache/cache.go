package cache

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + os.Getenv("REDIS_PORT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if rdb.Ping().Val() == "" {
		log.Fatal("CACHE ERROR")
	}

	return rdb
}