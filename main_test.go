package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test /user endpoint

func TestGetUser(t *testing.T) {
	// Setup
	// ...
	router := SetupRouter()

	// Execute
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/user", nil)
	router.ServeHTTP(res, req)

	// Verify
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))
	assert.Equal(t, `{"name":"John Smith"}`, res.Body.String())
}
