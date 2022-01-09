package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var todoList []string

func serverWebsocket() {
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error", err)
		}

		defer conn.Close()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error", err)
			}
			fmt.Printf("time %.2fs \n", time.Since(start).Seconds())
			_ = conn.WriteMessage(mt, message)
		}
	})
}

func serverStatic() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})

	http.ListenAndServe(":8080", nil)
}

func main() {
	serverWebsocket()
	serverStatic()
}
