package controller

import (
	"echo-mongodb-api/config"
	"echo-mongodb-api/module/database"
	"echo-mongodb-api/util"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// PlayerSuite
type PlayerSuite struct {
	suite.Suite
}

var token string

// Setup
func (suite PlayerSuite) SetupSuite() {
	config.Init()
	database.Connect()
	util.ClearDB()

	// Create fake player
	util.HelperPlayerCreateFake()

	// Generate Token
	data := map[string]interface{}{
		"id": util.PlayerID,
	}
	token = util.GenerateUserToken(data)

}

// TearDown
func (suite PlayerSuite) TearDownSuite() {
	util.ClearDB()
}

// TestGetProfileWithToken
func (suite *PlayerSuite) TestGetProfileWithToken() {
	var response util.Response

	// Create Context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodGet, "/player/my-profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)

	// Call MyProfile
	MyProfile(c)

	// Parse
	json.Unmarshal([]byte(responseRecorder.Body.Bytes()), &response)

	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
	assert.Equal(suite.T(), "Success!", response["message"])
}

// Run PlayerSuite
func TestPlayerSuite(t *testing.T) {
	suite.Run(t, new(PlayerSuite))
}
