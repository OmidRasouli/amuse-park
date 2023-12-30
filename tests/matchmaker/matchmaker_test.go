package tests

import (
	"log"
	"testing"

	"github.com/OmidRasouli/amuse-park/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
}

var (
	players []server.Player
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (suite *testSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func (suite *testSuite) TearDownSuite() {
}

func (suite *testSuite) TestMatchMaker() {
	log.Printf("=====================================")
	log.Printf("🧪 Matchmaking tests started")
	players = []server.Player{
		{
			ID:       "player1",
			Username: "Player1",
		}, {
			ID:       "player2",
			Username: "Player2",
		}, {
			ID:       "player3",
			Username: "Player3",
		}, {
			ID:       "player4",
			Username: "Player4",
		}, {
			ID:       "player5",
			Username: "Player5",
		}, {
			ID:       "player6",
			Username: "Player6",
		}, {
			ID:       "player7",
			Username: "Player7",
		}, {
			ID:       "player8",
			Username: "Player8",
		}, {
			ID:       "player9",
			Username: "Player9",
		},
	}
	inGameQueue, waitingQueue := server.StartGame(players...)
	assert.Equal(suite.T(), len(inGameQueue), len(players)-1)
	assert.Equal(suite.T(), len(waitingQueue), len(players)%2)
	assert.Equal(suite.T(), inGameQueue[0].Username, players[0].Username)
	assert.Equal(suite.T(), inGameQueue[1].Username, players[1].Username)
	assert.Equal(suite.T(), inGameQueue[2].Username, players[2].Username)
	assert.Equal(suite.T(), inGameQueue[3].Username, players[3].Username)
	assert.Equal(suite.T(), inGameQueue[4].Username, players[4].Username)
	assert.Equal(suite.T(), inGameQueue[5].Username, players[5].Username)
	assert.Equal(suite.T(), inGameQueue[6].Username, players[6].Username)
	assert.Equal(suite.T(), inGameQueue[7].Username, players[7].Username)
	assert.Equal(suite.T(), waitingQueue[0].Username, players[8].Username)

	log.Printf("✅ Matchmaking tests passed")

}
