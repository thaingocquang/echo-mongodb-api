package test

import (
	"echo-mongodb-api/model"
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

// LoginSuite
type LoginSuite struct {
	suite.Suite
	e *echo.Echo
}

// Setup
func (suite *LoginSuite) SetupSuite() {
	suite.e = testutil.InitServer()
	testutil.HelperPlayerCreateFake()
}

// TearDown
func (suite *LoginSuite) TearDownSuite() {
	testutil.ClearDB()
}

// TestLoginSuccess
func (suite *LoginSuite) TestLoginSuccess() {
	var (
		body = model.LoginBody{
			Email:    "tnq@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	// request & response
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// run
	suite.e.ServeHTTP(rec, req)

	// parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)

	// test
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// TestLoginEmailRequired ...
func (suite *LoginSuite) TestLoginEmailRequired() {
	var (
		body = model.LoginBody{
			Email:    "",
			Password: "123456",
		}
		response util.Response
	)

	// request & response
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// run
	suite.e.ServeHTTP(rec, req)

	// parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)

	// test
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), "email: email is required.", response["message"])
}

// TestLoginInvalidEmail ...
func (suite *LoginSuite) TestLoginInvalidEmail() {
	var (
		body = model.LoginBody{
			Email:    "invalidemail",
			Password: "123456",
		}
		response util.Response
	)

	// request & response
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// run
	suite.e.ServeHTTP(rec, req)

	// parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)

	// test
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), "email: email not valid.", response["message"])
}

// Run LoginSuite
func TestLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginSuite))
}
