package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader with custom CheckOrigin to allow connections from any origin
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define a struct to match the expected JSON structure
type Message struct {
	Message string `json:"message"`
}

// HandleConnections handles incoming WebSocket connections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial HTTP connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer ws.Close()

	// Listening for messages from client
	for {
		var msg Message
		// Read JSON message into the `msg` struct
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break // Exit loop if there’s an error
		}
		fmt.Println("Received Message:", msg.Message)

		// Optionally, you can send a response back to the client
		response := fmt.Sprintf("Message received: %s", msg.Message)
		err = ws.WriteMessage(websocket.TextMessage, []byte(response))
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", HandleConnections)
	fmt.Println("Starting server at port", 9090)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
