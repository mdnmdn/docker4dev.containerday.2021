package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
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

func handleIndexTemplate(response http.ResponseWriter, request *http.Request) {
	tmplt := template.New("index.html")
	tmplt, err := tmplt.ParseFiles("static/index.html")

	if err != nil {
		fmt.Println("Error: ", err)
	}
	ctx := context.Background()

	vars := struct {
		Counter int
	}{
		int(rdb.Incr(ctx, "counter").Val()),
	}
	tmplt.Execute(response, vars)
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				logger.Println(r.Method, r.URL.Path, r.RemoteAddr)
				//logger.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
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

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static/assets"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.HandleFunc("/", handleIndexTemplate)

	router.HandleFunc("/api", handleRequest)
	router.HandleFunc("/api/redis/incr", handleRequestRedisIncr)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), logging(logger)(router))
}
