// handlers/websocket.go

package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Upgrader to upgrade HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

// Keep track of all WebSocket clients.
var clients = make(map[*websocket.Conn]bool)
var clientsMutex sync.Mutex

// WebSocketHandler handles the WebSocket requests.
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to websocket:", err)
		return
	}

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	defer func() {
		clientsMutex.Lock()
		delete(clients, conn)
		clientsMutex.Unlock()
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received message: %s", message)

		// Respond with "catch" if message is "button click"
		if string(message) == "button click" {
			err = conn.WriteMessage(websocket.TextMessage, []byte("catch"))
			if err != nil {
				log.Println("Error sending message:", err)
				break
			}
		}
	}
}
