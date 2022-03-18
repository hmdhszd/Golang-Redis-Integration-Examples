## Get Program

Get a copy of the program:

```bash
  git clone https://github.com/hmdhszd/Golang-Redis-Integration-Examples.git
```

Go to the project directory:

```bash
  cd Golang-Redis-Integration-Examples
```



## Run Redis

First, run Redis in docker:

```bash
  docker run --name redis-test-instance -p 6369:6379 -d redis
```


## Run Golang Program

Then, install packages using go get:

```bash
  go mod init myapp

  go get github.com/go-redis/redis

  go get github.com/gomodule/redigo/redis
````

## 1- Check connection to Redis

```bash
  go run 1-\ check-redis-connection.go
````

The output should look like this:

```bash
  PONG <nil>
````

## 2- Set and Get value in redis

```bash
  ▶ go run 2-\ Set\ and\ Get\ value\ in\ redis.go 
  
  Hamid
````

## 3- Set JSON value in redis

```bash
  ▶ go run 3-\ Set\ JSON\ value\ in\ redis.go 

  {"name":"Elliot","age":25}
````

## 4- Set value in redis with Do command

```bash
  ▶ go run 4-\ Set\ value\ in\ redis\ with\ Do\ command.go 

  Electric Ladyland added!
````

## 5- Get value from redis with Do command

```bash
  ▶ go run 5-\ Get\ value\ from\ redis\ with\ Do\ command.go

  Electric Ladyland by Jimi Hendrix: £4.95 [8 likes]
````
