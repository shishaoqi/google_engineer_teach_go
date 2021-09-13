package redisConn

import (
	"github.com/go-redis/redis/v7"
	"log"
	"shishaoGo/redis/RedisInAction/config"
)

func ConnectRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:	config.Addr,
		Password: config.Password,
		DB:		config.DB,
	})

	if _, err := conn.Ping().Result(); err != nil {
		log.Fatal("Connect to redis client faild, err: %v\n", err)
	}
	return conn
}
