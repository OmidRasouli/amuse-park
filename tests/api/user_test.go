package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OmidRasouli/amuse-park/database"
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/OmidRasouli/amuse-park/server"
	"github.com/OmidRasouli/amuse-park/statics"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
}

type RegisterResult struct {
	Account server.UserAccount `json:"account"`
	Token   string             `json:"token"`
}

var (
	testUserAccount server.UserAccount
	token           string
)

func router() *gin.Engine {
	router := gin.Default()

	publicRoutes := router.Group("/")
	publicRoutes.POST("/register", server.Register)
	publicRoutes.POST("/update-profile", server.UpdateProfile)

	return router
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+token)
	}
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (suite *testSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	statics.Read()
	router()
	database.Initialize(&models.Account{}, &models.Authentication{}, &models.Profile{})
}

func (suite *testSuite) TearDownSuite() {
	migrator := database.Migrator()
	migrator.DropTable(&models.Account{})
	migrator.DropTable(&models.Authentication{})
	migrator.DropTable(&models.Profile{})
}

func (suite *testSuite) TestRegister() {
	testUser := server.UserAccount{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "testpassword",
	}
	response := makeRequest("POST", "/register", testUser, false)
	assert.Equal(suite.T(), http.StatusOK, response.Code)

	var registerResult RegisterResult
	err := json.Unmarshal(response.Body.Bytes(), &registerResult)
	assert.NoError(suite.T(), err, "Failed to unmarshal response JSON")

	token = registerResult.Token

	assert.Equal(suite.T(), testUser.Username, registerResult.Account.Username)
	testUserAccount.UserID = registerResult.Account.UserID
}

func (suite *testSuite) TestUpdateProfile() {
	userID := uuid.MustParse(testUserAccount.UserID)
	updatedProfile := models.Profile{
		ID:          userID,
		Level:       2,
		DisplayName: "UpdatedDisplayName",
		AvatarURL:   "https://example.com/avatar.jpg",
		Location:    "UpdatedLocation",
		TimeZone:    "UpdatedTimeZone",
		State:       "inactive",
		Email:       "updated_email@example.com",
	}

	response := makeRequest("POST", "/update-profile", updatedProfile, true)
	assert.Equal(suite.T(), http.StatusOK, response.Code)

	var updateProfile models.Profile
	err := json.Unmarshal(response.Body.Bytes(), &updateProfile)
	assert.NoError(suite.T(), err, "Failed to unmarshal response JSON")

	assert.Equal(suite.T(), updatedProfile.Level, updateProfile.Level)
	assert.Equal(suite.T(), updatedProfile.DisplayName, updateProfile.DisplayName)
	assert.Equal(suite.T(), updatedProfile.AvatarURL, updateProfile.AvatarURL)
	assert.Equal(suite.T(), updatedProfile.Location, updateProfile.Location)
	assert.Equal(suite.T(), updatedProfile.TimeZone, updateProfile.TimeZone)
	assert.Equal(suite.T(), updatedProfile.State, updateProfile.State)
	assert.Equal(suite.T(), updatedProfile.Email, updateProfile.Email)
}
