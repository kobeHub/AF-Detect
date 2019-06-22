package main
//
//import (
//	"github.com/gomodule/redigo/redis"
//)
//
//// The connect pool
//var redisPool *redis.Pool
//
//// Get connect instance to docker redis 
//func RedisPolling() *redis.Pool {
//	return &redis.Pool{
//		MaxIdle: 3,
//		MaxActive: 5,
//		Dial: func() (redis.Conn, error) {
//			c, err := redis.Dial("tcp", "127.0.0.1:6379")
//			if err != nil {
//				return nil, err
//			}
//
//			return c, err
//		},
//	}
//}
