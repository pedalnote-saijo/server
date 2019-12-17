package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/handle", http.HandlerFunc(testHandle))
	mux.Handle("/test/", http.HandlerFunc(testfunc))
	mux.Handle("/", http.HandlerFunc(indexHandler))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func testfunc(w http.ResponseWriter, r *http.Request) {
	// /handle/... is not routed
	// /handle is not routed
	log.Printf("path: %s", r.URL.Path)
}

func testHandle(w http.ResponseWriter, r *http.Request) {
	// /handle/test is not routed
	// /handle is routed
	log.Printf("path: %s", r.URL.Path)
}
