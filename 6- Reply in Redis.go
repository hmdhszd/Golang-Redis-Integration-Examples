package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Issue a HGET command to retrieve the title for a specific album,
	// and use the Str() helper method to convert the reply to a string.
	title, err := redis.String(conn.Do("HGET", "album:1", "title"))
	if err != nil {
		log.Fatal(err)
	}

	// Similarly, get the artist and convert it to a string.
	artist, err := redis.String(conn.Do("HGET", "album:1", "artist"))
	if err != nil {
		log.Fatal(err)
	}

	// And the price as a float64...
	price, err := redis.Float64(conn.Do("HGET", "album:1", "price"))
	if err != nil {
		log.Fatal(err)
	}

	// And the number of likes as an integer.
	likes, err := redis.Int(conn.Do("HGET", "album:1", "likes"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)
}
