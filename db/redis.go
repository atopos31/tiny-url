package db

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func init() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	db := os.Getenv("REDIS_DB")
	password := os.Getenv("REDIS_PASSWORD")
	dbnum, err := strconv.Atoi(db)
	if err != nil {
		panic(err)
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       dbnum, // use default DB
	})
}
