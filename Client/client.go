package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	// Connect to the WebSocket server
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:9090/ws", nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer c.Close()

	// Start a goroutine to handle incoming messages
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			fmt.Println("Received Message:", string(message))
		}
	}()

	// Main loop to send JSON messages from console input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message to send (or type 'exit' to quit): ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1] // Trim newline

		if text == "exit" {
			fmt.Println("Closing connection...")
			break
		}

		// Create a JSON message with a `message` field
		msg := map[string]string{"message": text}
		jsonMessage, err := json.Marshal(msg)
		if err != nil {
			log.Println("JSON marshal error:", err)
			continue
		}

		// Send JSON-formatted message to the server
		err = c.WriteMessage(websocket.TextMessage, jsonMessage)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
