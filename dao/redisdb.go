package dao

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var redisDB *redis.Pool

func InitRedis() {
	network := viper.GetString("redis.network")
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")

	redisDB = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(network,
				address,
				redis.DialPassword(password))
		},
	}

}

func GetRedisDB() *redis.Pool {
	return redisDB
}

func GetSessionStore() (sessions.RedisStore, error) {
	network := viper.GetString("redis.network")
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	size := viper.GetInt("redis.size")
	key := viper.GetString("redis.key")

	return sessions.NewRedisStore(size,
		network,
		address,
		password,
		[]byte(key))
}
