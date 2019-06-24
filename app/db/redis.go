package db

import (
	"sync"

	"github.com/go-redis/redis"
)

var gate = &sync.Mutex{}

var conn *redis.Client

// RedisClient : singleton client for redis connection
func RedisClient() *redis.Client {
	gate.Lock()
	defer gate.Unlock()

	if conn == nil {
		conn = redis.NewClient(&redis.Options{
			Addr:     "0.0.0.0:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}

	return conn
}
