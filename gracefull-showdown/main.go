package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := ":3000"
	server := &http.Server{Addr: port}

	log.Println("Server start at port: ", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Ola! segue sua resposta."))
	})

	go func() {
		fmt.Println("server is Running ...")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatal("Could not up the server")
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("shutting down server ...")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Could not gracefully shutdown the server", err)
	}

	fmt.Println("server Stopped")
}
