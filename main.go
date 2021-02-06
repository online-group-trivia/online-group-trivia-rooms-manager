package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path != "" {
		fmt.Fprintf(w, "Hi there, I love %s!", path)
	} else {
		fmt.Fprintf(w, "Hi there!")
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
	http.HandleFunc("/", handler)
	fmt.Printf("listening on %v\n", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux)))
}
