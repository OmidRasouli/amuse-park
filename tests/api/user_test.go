package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/OmidRasouli/amuse-park/database"
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/OmidRasouli/amuse-park/server"
	"github.com/OmidRasouli/amuse-park/statics"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type RegisterResult struct {
	Account server.UserAccount `json:"account"`
	Token   string             `json:"token"`
}

func TestRegister(t *testing.T) {
	testUser := server.UserAccount{
		UserID:   "test-user-id",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "testpassword",
	}
	response := makeRequest("POST", "/register", testUser, false)
	assert.Equal(t, http.StatusOK, response.Code)

	var registerResult RegisterResult
	err := json.Unmarshal(response.Body.Bytes(), &registerResult)
	assert.NoError(t, err, "Failed to unmarshal response JSON")

	log.Printf("the username is: %v", registerResult.Account.Username)
	assert.Equal(t, testUser.Username, registerResult.Account.Username)
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+bearerToken())
	}
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func bearerToken() string {
	user := server.UserAccount{
		UserID:   "test-user-id",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "testpassword",
	}

	writer := makeRequest("POST", "/register", user, false)
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response["jwt"]
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	statics.Read()
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()

	publicRoutes := router.Group("/")
	publicRoutes.POST("/register", server.Register)

	return router
}

func setup() {
	database.Initialize(&models.Account{}, &models.Authentication{}, &models.Profile{})
}

func teardown() {
	migrator := database.Migrator()
	migrator.DropTable(&models.Account{})
	migrator.DropTable(&models.Authentication{})
	migrator.DropTable(&models.Profile{})
}
