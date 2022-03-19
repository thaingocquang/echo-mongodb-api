package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// RunAndAssertHTTPOk ...
func RunAndAssertHTTPOk(e *echo.Echo, req *http.Request, t *testing.T) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	return rec
}

// RunAndAssertHTTPBadRequest ...
func RunAndAssertHTTPBadRequest(e *echo.Echo, req *http.Request, t *testing.T) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	return rec
}

// RunAndAssertHTTPBadRequest ...
func RunAndAssertHTTPUnauthorized(e *echo.Echo, req *http.Request, t *testing.T) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	return rec
}
