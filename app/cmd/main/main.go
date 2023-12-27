package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Maze struct {
	Size    int            `json:"size"`
	Content map[string]int `json:"content"`
}

func main() {
	http.HandleFunc("/", handleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Create a sample maze JSON
	maze := Maze{
		Size: 64,
		Content: map[string]int{
			"0.0": 0,
			"0.1": 1,
			// Add more maze content here
		},
	}

	// Convert maze to JSON
	mazeJSON, err := json.Marshal(maze)
	if err != nil {
		log.Println(err)
		return
	}

	// Send the maze JSON to the client
	if err := conn.WriteMessage(websocket.TextMessage, mazeJSON); err != nil {
		log.Println(err)
		return
	}

	log.Println("Maze sent to client")

	// Handle messages from the client (if needed)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
	}
}
