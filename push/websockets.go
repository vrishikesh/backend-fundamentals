package push

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %s", err)
		return
	}

	name := ws.RemoteAddr().String()
	clients[name] = ws

	m := fmt.Sprintf("%s %s", name, "connected!")
	log.Println(m)

	for _, conn := range clients {
		err = conn.WriteMessage(1, []byte(m))
		if err != nil {
			log.Printf("Error writing message: %s", err)
			return
		}
	}

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %s", err)
			return
		}

		name := conn.RemoteAddr().String()
		m := fmt.Sprintf("%s: %s", name, p)
		log.Printf("[Message] %s", m)

		for _, c := range clients {
			if err := c.WriteMessage(messageType, []byte(m)); err != nil {
				log.Printf("Error writing message: %s", err)
				return
			}
		}
	}
}

func Websockets() {
	fmt.Println("Go Websockets")
	setupRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
