package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var redisClient = RedisClient{
	client: nil,
}
var ctx = context.Background()

func echo(w http.ResponseWriter, r *http.Request) {
	rdb := GetClient(&redisClient)
	err := rdb.Publish(ctx, "mychannel1", "payload").Err()
	if err != nil {
		panic(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	httpPort := 9632
	http.HandleFunc("/", echo)
	fmt.Printf("listening on %v\n", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux)))
}
