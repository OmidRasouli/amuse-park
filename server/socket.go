package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type ServerMessage struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		if messageType == websocket.TextMessage {
			fmt.Printf("Received text message: %s\n", string(p))
			var serverMessage ServerMessage
			err := json.Unmarshal(p, &serverMessage)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
			if serverMessage.Type == "start_game" {
				player := Player{
					ID:       serverMessage.Data["ID"],
					Username: serverMessage.Data["Username"],
				}
				StartGame(player)
				fmt.Printf("data: %v -> %T\n", serverMessage.Data, serverMessage.Data)
			}
		} else if messageType == websocket.BinaryMessage {
			fmt.Printf("Received binary message: %v\n", p)
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func StartSocket() {
	http.HandleFunc("/ws", handleConnections)

	port := 8081
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
