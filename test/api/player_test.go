package test

import (
	"echo-mongodb-api/config"
	"echo-mongodb-api/module/database"
	"echo-mongodb-api/route"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"echo-mongodb-api/util"
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
	// table driven test
	tests := []struct {
		token          string
		wantMessage    string
		wantStatusCode int
	}{
		{
			token:          token,
			wantMessage:    "success",
			wantStatusCode: http.StatusOK,
		},
		{
			token:          "asasdasd", // wrong token
			wantMessage:    "invalid or expired jwt",
			wantStatusCode: http.StatusUnauthorized,
		},
	}

	// init server
	e := echo.New()
	route.Route(e)

	for _, test := range tests {
		// response
		var response util.Response

		// request & response
		req := httptest.NewRequest(http.MethodGet, "/players/my-profile", nil)
		req.Header.Set("Authorization", "Bearer "+test.token)
		rec := httptest.NewRecorder()

		// run
		e.ServeHTTP(rec, req)

		// parse
		json.Unmarshal([]byte(rec.Body.Bytes()), &response)

		// test
		assert.Equal(suite.T(), test.wantStatusCode, rec.Code)
		assert.Equal(suite.T(), test.wantMessage, response["message"])
	}
}

// Run PlayerSuite
func TestPlayerSuite(t *testing.T) {
	suite.Run(t, new(PlayerSuite))
}
