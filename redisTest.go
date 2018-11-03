package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("This is a program for go to use go_redis.")

	//connect
	cl := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	cl.WrapProcess(func(old func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			fmt.Printf("starting process:<%s>\n", cmd)
			err := old(cmd)
			fmt.Printf("finished process:<%s>\n", cmd)
			return err
		}
	})

	//get
	Get := func(client *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("get", key)
		client.Process(cmd)
		return cmd
	}

	//set
	Set := func(client *redis.Client, key string, value string) *redis.StringCmd {
		cmd := redis.NewStringCmd("set", key, value)
		client.Process(cmd)
		return cmd
	}

	//use
	_, errSet := Set(cl, "myKey", "myValue").Result()
	if errSet != nil {
		fmt.Println("redis set failed.", errSet)
	}
	value, errGet := Get(cl, "myKey").Result()
	if errGet != nil {
		fmt.Println("redis get failed.", errGet)
	} else {
		fmt.Println("The key value is", value)
	}

	//something else
	cmd := redis.NewStringCmd("set", "myKey1", "123", "ex", "100")
	cl.Process(cmd)

	//get expire
	cmd1 := redis.NewIntCmd("ttl", "myKey1")
	cl.Process(cmd1)
	expire, errEx := cmd1.Result()
	if errEx != nil {
		fmt.Println("ttl failed.", errEx)
	} else {
		fmt.Println("expire of key is", expire)
	}

}
