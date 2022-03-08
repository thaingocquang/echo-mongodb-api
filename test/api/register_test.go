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

// RegisterSuite
type RegisterSuite struct {
	suite.Suite
	e *echo.Echo
}

// Setup
func (suite *RegisterSuite) SetupSuite() {
	suite.e = testutil.InitServer()
}

// TearDown
func (suite *RegisterSuite) TearDownSuite() {
	testutil.ClearDB()
}

// TestRegisterSuccess ...
func (suite *RegisterSuite) TestRegisterSuccess() {
	var (
		body = model.PlayerCreateBody{
			Name:     "quang",
			Email:    "quang@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	// request & response
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", testutil.BodyToReader(body))
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

func (suite *RegisterSuite) TestRegisterEmailRequired() {
	var (
		body = model.PlayerCreateBody{
			Name:     "quang",
			Email:    "",
			Password: "123456",
		}
		response util.Response
	)

	// request & response
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", testutil.BodyToReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// run
	suite.e.ServeHTTP(rec, req)

	// parse
	json.Unmarshal([]byte(rec.Body.Bytes()), &response)

	// test
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

// Run RegisterSuite
func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
}
