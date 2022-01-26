package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "localhost:6369")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// Send our command across the connection. The first parameter to
	// Do() is always the name of the Redis command (in this example
	// HMSET), optionally followed by any necessary arguments (in this
	// example the key, followed by the various hash fields and values).

	_, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Electric Ladyland added!")
}
