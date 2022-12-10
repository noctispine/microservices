package cache

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if rdb.Ping().Val() == "" {
		log.Fatal("CACHE ERROR: ", rdb.Ping())
	}

	return rdb
}