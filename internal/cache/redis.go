package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)
var Client* redis.Client;

var Ctx = context.Background();

func InitRedis(){
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",  
		DB: 0, 
	})
	
	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}
}