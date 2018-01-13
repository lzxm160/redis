package main

import (
	"fmt"
	redis "github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "r-bp19e097f1b37414.redis.rds.aliyuncs.com:6379",
		Password: "Miner7032018", // no password set
		DB:       0,              // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// err = client.Set("key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := client.Get("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	val, err := client.HGetAll("eth:{xx}"))
	if err == redis.Nil {
		fmt.Println("eth:{xx} does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("eth:{xx} ", val)
	}
	val2, err := client.HGetAll("eth:{xx}:xx"))
	if err == redis.Nil {
		fmt.Println("eth:{xx}:xx does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("eth:{xx}:xx ", val2)
	}
}
