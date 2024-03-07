package gredis

import "github.com/gomodule/redigo/redis"

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:
	}

}