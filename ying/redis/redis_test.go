package redis

import (
	"github.com/go-redis/redis/v7"
	log "github.com/pion/ion-log"
	"testing"
)

//
// GetRedisClient
//  @Description: 创建redis客户端
//
func GetRedisClient() *redis.Client {
	log.Infof("Connecting to Redis-server...")
	err := Init()
	if err != nil {
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addresses[0],
		Password: config.Redis.Pwd,
		DB:       0,
	})

	err = client.Ping().Err()
	if err != nil {
		log.Errorf("Ping Redis-Server Failed! Reason: %v", err)
		return nil
	}

	log.Infof("Connect to Redis-server Success!")
	return client
}

func TestConnect(t *testing.T) {
	GetRedisClient()
}
