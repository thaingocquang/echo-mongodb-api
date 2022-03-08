package test

import (
	testutil "echo-mongodb-api/test_util"
	"echo-mongodb-api/util"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// GetProfileSuite
type GetProfileSuite struct {
	suite.Suite
	e *echo.Echo
}

var token string

// Setup
func (suite *GetProfileSuite) SetupSuite() {
	suite.e = testutil.InitServer()

	// Create fake player
	testutil.HelperPlayerCreateFake()

	// Generate Token
	data := map[string]interface{}{
		"id": testutil.PlayerID,
	}
	token = util.GenerateUserToken(data)
}

// TearDown
func (suite *GetProfileSuite) TearDownSuite() {
	testutil.ClearDB()
}

// TestGetProfileSuccess ...
func (suite *GetProfileSuite) TestGetProfileSuccess() {
	// response
	var response util.Response

	// request & response
	req := httptest.NewRequest(http.MethodGet, "/players/my-profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	// run
	suite.e.ServeHTTP(rec, req)

	// parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)

	// test
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// TestGetProfileInvalidToken ...
func (suite *GetProfileSuite) TestGetProfileInvalidToken() {
	// response
	var response util.Response
	invalidToken := "asdasd"

	// request & response
	req := httptest.NewRequest(http.MethodGet, "/players/my-profile", nil)
	req.Header.Set("Authorization", "Bearer "+invalidToken)
	rec := httptest.NewRecorder()

	// run
	suite.e.ServeHTTP(rec, req)

	// parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)

	// test
	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

// Run GetProfileSuite
func TestGetProfileSuite(t *testing.T) {
	suite.Run(t, new(GetProfileSuite))
}
