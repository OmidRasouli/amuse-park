package server

import (
	"fmt"

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

		createRoom(waitingQueue[:maxPlayerInRoom]...)
		inGameQueue = append(inGameQueue, waitingQueue[:maxPlayerInRoom]...)
		waitingQueue = waitingQueue[maxPlayerInRoom:]
	}
}

func createRoom(players ...Player) {
	newRoom := Room{
		ID:      uuid.NewString(),
		Players: players,
	}
	rooms = append(rooms, newRoom)
	matchmaking(newRoom)
}

func matchmaking(room Room) {
	startMatch(room)
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
