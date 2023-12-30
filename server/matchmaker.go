package server

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Player struct {
	ID       string
	Username string
	Status   string
}

type Room struct {
	ID      string
	Players []Player
}

var (
	waitingQueue    []Player
	inGameQueue     []Player
	rooms           []Room
	maxPlayerInRoom int = 2
)

func StartGameRequest(player Player) {
	waitingQueue = append(waitingQueue, player)
	if len(waitingQueue) == maxPlayerInRoom {

		log.Printf("players in waiting queue: %v", waitingQueue[:maxPlayerInRoom])
		createRoom(waitingQueue[:maxPlayerInRoom]...)
		log.Printf("players in waiting queue again: %v", waitingQueue)
		inGameQueue = append(inGameQueue, waitingQueue[:maxPlayerInRoom]...)
		log.Printf("playersin in game queue: %v", inGameQueue)
		log.Printf("playersin waiting queue before the end: %v", waitingQueue)
		waitingQueue = waitingQueue[maxPlayerInRoom:]
		log.Printf("playersin waiting queue at the end: %v", waitingQueue)
	}
}

func createRoom(players ...Player) {
	newRoom := Room{
		ID:      uuid.NewString(),
		Players: players,
	}
	rooms = append(rooms, newRoom)
	fmt.Println("Room created with players:", players)
	matchmaking(newRoom)
}

func matchmaking(room Room) {
	startMatch(room)
	fmt.Printf("remainig players in queue: %v", waitingQueue)
}

func startMatch(room Room) {
	var playersName string
	for i := 0; i < maxPlayerInRoom; i++ {
		playersName = fmt.Sprintf("player %v: %v", (i + 1), room.Players[i].Username)
	}
	fmt.Printf("game started with room id: %v and players: %v\n", room.ID, playersName)
}

func StartGame(players ...Player) ([]Player, []Player) {
	for _, player := range players {
		StartGameRequest(player)
	}
	return inGameQueue, waitingQueue
}
