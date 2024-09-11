package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow CORS
	},
}

// handles incoming WebSocket connections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket:", err)
		return
	}
	defer ws.Close()

	fmt.Println("A user connected")

	// Send random data every second
	for {
		randomLine := generateRandomData()
		err := ws.WriteMessage(websocket.TextMessage, []byte(randomLine))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("A user disconnected")
}
