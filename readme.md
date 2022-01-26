

// Run Redis

▶ docker run --name redis-test-instance -p 6379:6379 -d redis




// get requirements and run the program

▶ go mod init myapp

▶ go get github.com/go-redis/redis

▶ go get github.com/gomodule/redigo/redis




// Run the program

▶ go run 1-\ check-redis-connection.go

PONG <nil>

