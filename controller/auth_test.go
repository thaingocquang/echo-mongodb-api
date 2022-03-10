package controller_test

import (
	"echo-mongodb-api/model"
	testutil "echo-mongodb-api/test_util"
	"echo-mongodb-api/util"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// authVerbs ...
var authVerbs = map[string]*testutil.Verb{
	"login":    {Method: http.MethodPost, Url: "/auth/login"},
	"register": {Method: http.MethodPost, Url: "/auth/register"},
}

// LoginSuite ...
type LoginSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *LoginSuite) SetupSuite() {
	suite.e = testutil.InitServer()
	testutil.HelperPlayerCreateFake()
}

// TearDownSuite ...
func (suite *LoginSuite) TearDownSuite() {
}

// TestLoginSuccess ...
func (suite *LoginSuite) TestLoginSuccess() {
	var (
		body = model.LoginBody{
			Email:    "tnq@gmail.com",
			Password: "123456",
		}
		response  util.Response
		loginVerb = authVerbs["login"]
	)

	// request
	req, _ := http.NewRequest(loginVerb.Method, loginVerb.Url, testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testutil.RunAndAssertHTTPOk(suite.e, req, suite.T())

	// parse response
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// TestLoginEmailRequired ...
func (suite *LoginSuite) TestLoginEmailRequired() {
	var (
		body = model.LoginBody{
			Email:    "",
			Password: "123456",
		}
		response  util.Response
		loginVerb = authVerbs["login"]
	)

	// request
	req, _ := http.NewRequest(loginVerb.Method, loginVerb.Url, testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testutil.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// parse response
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)
	assert.Equal(suite.T(), "email: email is required.", response["message"])
}

// TestLoginInvalidEmail ...
func (suite *LoginSuite) TestLoginInvalidEmail() {
	var (
		body = model.LoginBody{
			Email:    "invalidemail",
			Password: "123456",
		}
		response  util.Response
		loginVerb = authVerbs["login"]
	)

	// request
	req, _ := http.NewRequest(loginVerb.Method, loginVerb.Url, testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testutil.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)
	assert.Equal(suite.T(), "email: email not valid.", response["message"])
}

// RegisterSuite ...
type RegisterSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *RegisterSuite) SetupSuite() {
	suite.e = testutil.InitServer()
}

// TearDownSuite ...
func (suite *RegisterSuite) TearDownSuite() {
}

// TestRegisterSuccess ...
func (suite *RegisterSuite) TestRegisterSuccess() {
	var (
		body = model.PlayerCreateBody{
			Name:     "quang",
			Email:    "quang@gmail.com",
			Password: "123456",
		}
		response  util.Response
		loginVerb = authVerbs["register"]
	)

	// request
	req, _ := http.NewRequest(loginVerb.Method, loginVerb.Url, testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testutil.RunAndAssertHTTPOk(suite.e, req, suite.T())

	json.Unmarshal([]byte(rec.Body.Bytes()), &response)
	assert.NotEqual(suite.T(), nil, response["data"])
}

func (suite *RegisterSuite) TestRegisterEmailRequired() {
	var (
		body = model.PlayerCreateBody{
			Name:     "quang",
			Email:    "",
			Password: "123456",
		}
		response  util.Response
		loginVerb = authVerbs["register"]
	)

	// request
	req, _ := http.NewRequest(loginVerb.Method, loginVerb.Url, testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := testutil.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	json.Unmarshal([]byte(rec.Body.Bytes()), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

// TestAuthTestSuite ...
func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
	suite.Run(t, new(LoginSuite))
}
