package redis

import (
	"github.com/go-redis/redis/v7"
	log "github.com/pion/ion-log"
	"testing"
)

var client *redis.Client

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

func InitClient() {
	client = GetRedisClient()
}

func TestConnect(t *testing.T) {
	InitClient()
}

type Person struct {
	Name string
	Age  int
}

func TestSet(t *testing.T) {
	InitClient()
	err := client.Set("s1", "hello", 0).Err()
	if err != nil {
		log.Errorf("Set s1 Failed! Reason: %v", err)
		return
	}

	s1 := client.Get("s1").String()
	log.Infof("Get s1 Success! s1: %v", s1)

}

//
//  TestSetStruct
//  @Description: 测试set结构体,失败，无法set结构体
//
func TestSetStruct(t *testing.T) {
	InitClient()
	//  测试set结构体
	//  失败
	var p1 Person
	p1.Name = "郑晓颖"
	p1.Age = 18
	err := client.Set("s2", p1, 0).Err()
	if err != nil {
		log.Errorf("Set s2 Failed! Reason: %v", err)
		return
	}

	s2 := client.Get("s2").String()
	log.Infof("Get s2 Success! s1: %v", s2)
}

//
//  TestSetNX
//  @Description: 测试set nx
//  @param t
//
func TestSetNX(t *testing.T) {
	InitClient()
	//  测试nx
	success := client.SetNX("s1", "nx", 0).Val()
	if !success {
		log.Errorf("SetNX s1 Failed!")
	}

	s1 := client.Get("s1").String()
	log.Infof("Get s1 Success! s1: %v", s1)
}
