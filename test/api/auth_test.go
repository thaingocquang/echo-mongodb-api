package test

import (
	"bytes"
	"echo-mongodb-api/config"

	// "echo-mongodb-api/controller"

	// "echo-mongodb-api/controller"
	"echo-mongodb-api/model"
	"echo-mongodb-api/module/database"
	"echo-mongodb-api/route"
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

func (suite *RegisterSuite) TestRegister() {
	// table driven test
	tests := []struct {
		body        model.PlayerCreateBody
		wantMessage string
		statusCode  int
	}{
		{
			body:        model.PlayerCreateBody{Name: "quang", Email: "quang@gmai.com", Password: "123456"},
			wantMessage: "success",
			statusCode:  http.StatusOK,
		},
		{
			body:        model.PlayerCreateBody{Name: "q", Email: "quang@gmai.com", Password: "123456"},
			wantMessage: "name: name is length: 3 -> 30.",
			statusCode:  http.StatusBadRequest,
		},
	}

	// init server
	e := echo.New()
	route.Route(e)

	for _, test := range tests {
		// response
		var response util.Response

		// test body
		testBody := map[string]interface{}{
			"name":     test.body.Name,
			"email":    test.body.Email,
			"password": test.body.Password,
		}

		// marshal body
		body, err := json.Marshal(testBody)

		// marshal error
		if err != nil {
			suite.T().Error("Marshal body failed")
		}

		// request & response
		req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		// run
		e.ServeHTTP(rec, req)

		// Parse
		json.Unmarshal([]byte(rec.Body.Bytes()), &response)

		// Test
		assert.Equal(suite.T(), test.statusCode, rec.Code)
		assert.Equal(suite.T(), test.wantMessage, response["message"])
	}
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
	// table driven test
	tests := []struct {
		body           model.LoginBody
		wantMessage    string
		wantStatusCode int
	}{
		{
			body:           model.LoginBody{Email: "tnq@gmail.com", Password: "123456"},
			wantMessage:    "success",
			wantStatusCode: http.StatusOK,
		},
		{
			body:           model.LoginBody{Email: "", Password: "123456"},
			wantMessage:    "email: email is required.",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			body:           model.LoginBody{Email: "invalidemail", Password: "123456"},
			wantMessage:    "email: email not valid.",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			body:           model.LoginBody{Email: "tnq@gmail.com", Password: ""},
			wantMessage:    "password: password is required.",
			wantStatusCode: http.StatusBadRequest,
		},
	}

	// init server
	e := echo.New()
	route.Route(e)

	for _, test := range tests {
		// response
		var response util.Response

		// test body
		testBody := map[string]interface{}{
			"email":    test.body.Email,
			"password": test.body.Password,
		}

		// marshal body
		body, err := json.Marshal(testBody)

		// marshal error
		if err != nil {
			suite.T().Error("Marshal body failed")
		}

		// request & response
		req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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

// Run RegisterSuite & LoginSuite
func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
	suite.Run(t, new(LoginSuite))
}
