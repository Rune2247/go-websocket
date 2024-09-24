// main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"server/handlers" // Updated to reflect the new module name
)

func main() {
	// Route the /ws path to the WebSocket handler
	http.HandleFunc("/ws", handlers.WebSocketHandler)

	// Start the server on port 8080
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
