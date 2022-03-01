package controller

import (
	"echo-mongodb-api/config"
	"echo-mongodb-api/model"
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

// RegisterSuite
type RegisterSuite struct {
	suite.Suite
}

// Setup
func (suite RegisterSuite) SetupSuite() {
	config.Init()
	database.Connect()
	util.ClearDB()
}

// TearDown
func (suite RegisterSuite) TearDownSuite() {
	util.ClearDB()
}

//Test
func (suite *RegisterSuite) TestRegister() {
	var (
		playerRegisterTest = model.PlayerCreateBody{
			Name:     "thaingocquang",
			Email:    "quang@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	// Create Context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", playerRegisterTest)

	// Call Register
	Register(c)

	// Parse
	json.Unmarshal([]byte(responseRecorder.Body.Bytes()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
	assert.Equal(suite.T(), "Success!", response["message"])
}

// LoginSuite
type LoginSuite struct {
	suite.Suite
}

// Setup
func (suite LoginSuite) SetupSuite() {
	config.Init()
	database.Connect()
	util.ClearDB()
	util.HelperPlayerCreateFake()
}

// TearDown
func (suite LoginSuite) TearDownSuite() {
	util.ClearDB()
}

// Test
func (suite *LoginSuite) TestLogin() {
	var (
		playerLoginTest = model.LoginBody{
			Email:    "thaingocquang147@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	//Create Context
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	responseRecorder := httptest.NewRecorder()
	c := e.NewContext(req, responseRecorder)
	c.Set("body", playerLoginTest)

	// Call Login
	Login(c)

	//Parse
	json.Unmarshal([]byte(responseRecorder.Body.Bytes()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "Success!", response["message"])

}

// Run RegisterSuite & LoginSuite
func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
	suite.Run(t, new(LoginSuite))
}
