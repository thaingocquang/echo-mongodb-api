package controller_test

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

// playerVerbs ...
var playerVerbs = map[string]*testutil.Verb{
	"myprofile": {Method: http.MethodGet, Url: "/players/my-profile"},
}

// GetProfileSuite ...
type GetProfileSuite struct {
	suite.Suite
	e *echo.Echo
}

var token string

// SetupSuite ...
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

// TearDownSuite ...
func (suite *GetProfileSuite) TearDownSuite() {
}

// TestGetProfileSuccess ...
func (suite *GetProfileSuite) TestGetProfileSuccess() {
	// response
	var (
		response       util.Response
		getProfileVerb = playerVerbs["myprofile"]
	)

	// request
	req := httptest.NewRequest(getProfileVerb.Method, getProfileVerb.Url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// run
	rec := testutil.RunAndAssertHTTPOk(suite.e, req, suite.T())

	// parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// TestGetProfileInvalidToken ...
func (suite *GetProfileSuite) TestGetProfileInvalidToken() {
	// response
	var (
		response       util.Response
		invalidToken   = "asdasd"
		getProfileVerb = playerVerbs["myprofile"]
	)

	// request
	req := httptest.NewRequest(getProfileVerb.Method, getProfileVerb.Url, nil)
	req.Header.Set("Authorization", "Bearer "+invalidToken)

	// run
	rec := testutil.RunAndAssertHTTPUnauthorized(suite.e, req, suite.T())

	// // parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

// TestPlayerTestSuite ...
func TestPlayerTestSuite(t *testing.T) {
	suite.Run(t, new(GetProfileSuite))
}
