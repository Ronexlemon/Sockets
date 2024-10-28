package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)



func main(){
	c,_,err := websocket.DefaultDialer.Dial("ws://localhost:9090",nil)
	if err !=nil{
		log.Fatal(err)
	}
	defer c.Close()
	for {
		_,Message,err:=c.ReadMessage()
		if err!=nil{
			log.Fatal(err)
			}
			fmt.Println("Receive Message",Message)
	}
}