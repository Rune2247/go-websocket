// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Define an upgrader to upgrade HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity (you might want to restrict this in production)
		return true
	},
}

// Keep track of all active WebSocket clients.
var clients = make(map[*websocket.Conn]bool)
var clientsMutex sync.Mutex

// Handle WebSocket connections.
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to websocket:", err)
		return
	}

	// Add the new client to the list of active clients.
	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	defer func() {
		// Remove client on disconnect.
		clientsMutex.Lock()
		delete(clients, conn)
		clientsMutex.Unlock()
		conn.Close()
	}()

	// Listen for messages from the client.
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received message: %s", message)

		// If the message is "button click", respond with "catch".
		if string(message) == "button click" {
			err = conn.WriteMessage(websocket.TextMessage, []byte("catch"))
			if err != nil {
				log.Println("Error sending message:", err)
				break
			}
		}
	}
}

func main() {
	// Handle WebSocket connections at the "/ws" endpoint.
	http.HandleFunc("/ws", handleWebSocket)

	// Start the HTTP server.
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
