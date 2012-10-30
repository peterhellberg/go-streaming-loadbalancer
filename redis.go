package main

import "github.com/garyburd/redigo/redis"

var pool *redis.Pool

func new_redis_pool() {
	pool = &redis.Pool{
		MaxIdle: 3,
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.Dial("tcp", *redis_server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

func GetRedirectIP(c redis.Conn) (string, error) {
	return redis.String(c.Do("GET", *redis_key))
}
