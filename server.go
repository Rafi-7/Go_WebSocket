package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Websocket server started on : 9898")
	setupRoutes()
	err := http.ListenAndServe(":9898", nil)
	if err != nil {
		fmt.Println("ListenandServer error : ", err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrader error : ", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error : ", err)
			break
		}
		fmt.Printf("Received : %s\n", message)
	}
	
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}
