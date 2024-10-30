package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


type Message struct{
	Username string `json:"username"`
	Message string `json:"message"`
}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/ws",handleConnections)
	go handleMessages()
	fmt.Println("Server started on :9090")
	err:= http.ListenAndServe(":8080",nil)
	if err !=nil{
		log.Fatal("Failed to server",err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to chat room")
}

func handleConnections(w http.ResponseWriter, r *http.Request){
	conn,err := upgrader.Upgrade(w,r,nil)
	if err !=nil{
		log.Fatal("Failed to open a connection")
		return
	}
	defer conn.Close()

	clients[conn] =true

	for{
		var msg Message
		err:= conn.ReadJSON(&msg)
		if err !=nil{
			log.Fatal(err)
			delete(clients,conn)
			return
		}
		broadcast <-msg
	}


}