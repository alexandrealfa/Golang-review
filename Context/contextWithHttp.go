package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func getBananas(w http.ResponseWriter, r *http.Request) {
	log.Println("request started")

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)

	defer log.Println("request ended")
	defer cancel()

	select {
	case <-time.After(time.Second * 3):
		log.Println("request processed")
		if _, err := w.Write([]byte("Request has Processed")); err != nil {
			log.Fatal(err)
		}
	case <-ctx.Done():
		log.Println("request finished without process, client cancel")
		// http.Error(w, "Request hasn't Processed because client canceled", http.StatusRequestTimeout)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getBananas)

	if err := http.ListenAndServe(":8089", mux); err != nil {
		log.Fatal(err)
	}
}
