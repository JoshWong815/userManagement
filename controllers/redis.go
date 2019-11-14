package controllers

import (
	"fmt"
	"github.com/go-redis/redis"
)

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	err1 := client.Set("key", "value", 0).Err()
	if err1 != nil {
		panic(err1)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	//删除
	n, err := client.Del("key").Result()

	fmt.Println(n, "条记录已被删除")
	if err != nil {
		panic(err)
	}
}



func getRedisConnected() *redis.Client{
	Client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return Client
}

