package main

import (
	"fmt"
	redis "github.com/go-redis/redis"
	"strconv"
	"sync"
)

func ExampleClient_Watch(client *redis.Client) {
	var incr func(string) error

	// Transactionally increments key using GET and SET commands.
	incr = func(key string) error {
		err := client.Watch(func(tx *redis.Tx) error {
			n, err := tx.Get(key).Int64()
			if err != nil && err != redis.Nil {
				return err
			}

			_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
				pipe.Set(key, strconv.FormatInt(n+1, 10), 0)
				return nil
			})
			return err
		}, key)
		if err == redis.TxFailedErr {
			return incr(key)
		}
		return err
	}

	// var wg sync.WaitGroup
	// for i := 0; i < 100; i++ {
	// // 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()

	err := incr("counter3")
	if err != nil {
		panic(err)
	}
	// 	}()
	// }
	// wg.Wait()

	n, err := client.Get("counter3").Int64()
	fmt.Println(n, err)
	// Output: 100 <nil>
}
func main() {
	client := redis.NewClient(&redis.Options{
		// Addr:     "r-bp19e097f1b37414.redis.rds.aliyuncs.com:6379",
		// Password: "Miner7032018", // no password set
		Addr:     "127.0.0.1:6379",
		Password: "etcpool123",
		DB:       0, // use default DB
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
	pipe := client.Pipeline()
	sel := pipe.Select(2)
	fmt.Println("select ", sel.String())

	val1 := pipe.HGetAll("eth:{xx}")
	fmt.Println("eth:{xx} ", val1.String())

	val2 := pipe.HGetAll("eth:{xx}:xx")
	fmt.Println("eth:{xx}:xx ", val2.String())
	_, err = pipe.Exec()
	fmt.Println("eth:{xx} ", val1.String())
	fmt.Println("eth:{xx}:xx ", val2.String())
	fmt.Println(err)
	fmt.Println("----------------------------")
	ExampleClient_Watch(client)
}
