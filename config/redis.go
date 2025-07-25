package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis(){
	RedisClient = redis.NewClient(&redis.Options{
		Addr:	os.Getenv("REDIS_ADDR"),
		Password:	os.Getenv("REDIS_PASSWORD"),
		DB:	0,
	})

	// Test Connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Connected to Redis")
}