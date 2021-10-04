package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func handleRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "ok - ")
}

func handleRequestRedisIncr(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	val := rdb.Incr(ctx, "counter")
	//val := "wow!"
	fmt.Fprint(res, "go bare ", val)
}

func main() {
	port := os.Getenv("PORT")
	redisUrl := os.Getenv("REDIS_URL")

	if port == "" {
		port = "5000"
	}

	if redisUrl != "" {
		rdb = redis.NewClient(&redis.Options{
			Addr: redisUrl,
		})
	}

	fmt.Println("starting...")

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/redis/incr", handleRequestRedisIncr)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
