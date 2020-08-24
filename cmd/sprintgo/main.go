package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	mymux := http.NewServeMux()
	mymux.HandleFunc("/", myHandler)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mymux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

