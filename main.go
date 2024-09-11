package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type randomData struct {
	CLOCK     string
	YAW       string
	PITCH     string
	ROLL      string
	LATITUDE  string
	LONGITUDE string
	VOLTAGE   string
	PRESSURE  string
	ALTITUDE  string
}

// Upgrader to upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow CORS
	},
}

func randomGenerator() string {
	now := time.Now()
	CLOCK := fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())

	YAW := fmt.Sprintf("%.2f", rand.Float64()*360-180)
	PITCH := fmt.Sprintf("%.2f", rand.Float64()*360-180)
	ROLL := fmt.Sprintf("%.2f", rand.Float64()*360-180)

	lat := -7.773684
	long := 110.381798
	LATITUDE := fmt.Sprintf("%.6f", lat+rand.Float64()*0.0002-0.0001)
	LONGITUDE := fmt.Sprintf("%.6f", long+rand.Float64()*0.0002-0.0001)

	VOLTAGE := fmt.Sprintf("%.2f", rand.Float64()*12)
	PRESSURE := fmt.Sprintf("%.2f", rand.Float64()*100)
	ALTITUDE := fmt.Sprintf("%.2f", rand.Float64()*700)

	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s;",
		CLOCK, YAW, PITCH, ROLL, LATITUDE, LONGITUDE, VOLTAGE, PRESSURE, ALTITUDE)
}

// handleConnections handles incoming WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial HTTP connection to WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket:", err)
		return
	}
	defer ws.Close()

	fmt.Println("A user connected")

	// Send random data every second
	for {
		randomLine := randomGenerator()
		err := ws.WriteMessage(websocket.TextMessage, []byte(randomLine))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("A user disconnected")
}

func main() {
	// Define HTTP server and WebSocket route
	http.HandleFunc("/ws", handleConnections)

	// Start HTTP server
	port := ":5000"
	fmt.Println("Server started on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
