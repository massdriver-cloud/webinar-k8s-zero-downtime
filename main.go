package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// simulate a 50-500ms response time
	latency := rand.Intn(450) + 50
	time.Sleep(time.Duration(latency) * time.Millisecond)
	fmt.Fprintf(w, "Hello, World!!\n")
}

func main() {
	log.Println("Server starting up.")

	// add startup latency
	time.Sleep(5 * time.Second)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)

	// server := &http.Server{
	// 	Addr: ":8080",
	// }

	// http.HandleFunc("/", hello)

	// go func() {
	// 	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("HTTP server error: %v", err)
	// 	}
	// 	log.Println("Stopped serving new connections.")
	// }()

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// <-sigChan

	// shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	// defer shutdownRelease()

	// if err := server.Shutdown(shutdownCtx); err != nil {
	// 	log.Fatalf("HTTP shutdown error: %v", err)
	// }

	log.Println("Graceful shutdown complete.")
}
