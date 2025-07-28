package main

import (
	"log"
	"net/http"
	"prontoShare/server/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.WSHandler)
	log.Println("Signaling server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
