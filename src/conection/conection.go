package conection

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ConectionDbUser() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return *client
}
