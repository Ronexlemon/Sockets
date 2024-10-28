package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
func HandleConnections(w http.ResponseWriter, r *http.Request){
	ws,err:=upgrader.Upgrade(w,r,nil)
	if err !=nil{
		log.Fatal("failed o start",err)
	}
	defer ws.Close()
	for{
		var msg string
		err = ws.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
	}
	fmt.Println("Received Message",msg)
}

}