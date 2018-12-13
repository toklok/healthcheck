package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthConnection(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := setupRouter()

	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Errorf("Get heartbeat failed with error %d", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("/healthcheck is 200 OK %d.", resp.Code)
	}
}
