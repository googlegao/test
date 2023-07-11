package test

import "github.com/go-redis/redis"

func ConnRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Ping().Err()

	return rdb, err
}
