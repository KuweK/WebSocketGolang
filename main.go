package main

import (
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsapi(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Server claim message: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func wspage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/sockets.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", wspage)
	http.HandleFunc("/ws", wsapi)
	http.ListenAndServe("localhost:8080", nil)
}
