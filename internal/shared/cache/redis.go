package cache

import (
	"context"
	"log"

	"github.com/NIROOZbx/project-server/internal/shared/config"
	"github.com/redis/go-redis/v9"
)


var rdb *redis.Client
func InitRedis() {

	
	opt, err := redis.ParseURL(config.GetConfig().RedisURL)

	if err != nil {
		log.Fatalf("❌ Could not parse Redis URL: %v", err)
	}

	rdb=redis.NewClient(opt)



	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("❌ Could not connect to Redis: %v", err)
	}

	log.Println("✅ Redis connected successfully.")
}

func GetClient() *redis.Client {
	return rdb
}