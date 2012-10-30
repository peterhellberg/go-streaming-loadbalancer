package main

import "github.com/garyburd/redigo/redis"

var pool *redis.Pool

func createRedisPool() {
	pool = &redis.Pool{
		MaxIdle: 3,
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.Dial("tcp", *redisServer)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

func getRedirectIP(c redis.Conn) (string, error) {
	return redis.String(c.Do("GET", *redisKey))
}
