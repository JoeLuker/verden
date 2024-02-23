// db/redis.go
package db

import (
	"context"
	"fmt"
	"log"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/JoeLuker/verden/models"
	"os"
)

// RedisService struct to hold the Redis client
type RedisService struct {
	Client *redis.Client
}

var RedisNil = redis.Nil


// NewRedisService creates a new instance of RedisService
func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{Client: client}
}

// ConnectRedis establishes a connection to the Redis server
func ConnectRedis(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"), // Redis service name and port
		Password: "",           // no password set
		DB:       0,            // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		fmt.Println("Connected to Redis successfully")
	}

	return rdb
}
// In redis.go
func (s *RedisService) Set(ctx context.Context, key string, value string) error {
    return s.Client.Set(ctx, key, value, 0).Err()
}

func (s *RedisService) Get(ctx context.Context, key string) (string, error) {
    val, err := s.Client.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", nil // Key does not exist
    } else if err != nil {
        return "", err
    }
    return val, nil
}


func (service *RedisService) SaveToRedis(ctx context.Context, userData *models.UserData) error {
    userDataJSON, err := json.Marshal(userData)
    if err != nil {
        return err
    }

    return service.Client.Set(ctx, userData.UserID, userDataJSON, 0).Err()
}

