package main

import (
	"fmt"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6369",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Person{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}

	//Set value
	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	//Get value
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
