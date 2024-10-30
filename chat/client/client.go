package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:9090/ws", nil)
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer conn.Close()

	// Start a goroutine to continuously listen for incoming messages
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error receiving message:", err)
				return
			}
			fmt.Println("Received message:", string(message))
		}
	}()

	// Take the username input from the user
	fmt.Print("Enter your username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading username:", err)
	}
	username = strings.TrimSpace(username)

	
	for {
		fmt.Println("Enter message: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading message:", err)
		}
		text = strings.TrimSpace(text)

		
		msg := Message{
			Username: username,
			Message:  text,
		}
		jsonMessage, err := json.Marshal(msg)
		if err != nil {
			log.Println("Error encoding message to JSON:", err)
			continue
		}

		err = conn.WriteMessage(websocket.TextMessage, jsonMessage)
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}
}
