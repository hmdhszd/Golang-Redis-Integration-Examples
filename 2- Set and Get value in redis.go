package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6369",
		Password: "",
		DB:       0,
	})

	// Set value

	err := client.Set("Name", "Hamid", 0).Err()

	if err != nil {
		fmt.Println(err)
	}

	// Get value

	val, err := client.Get("Name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

}
