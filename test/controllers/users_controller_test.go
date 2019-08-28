package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jigintern/Foodmates-server/test/initialize"
	"github.com/stretchr/testify/assert"
)

func TestReadUserSucceed(t *testing.T) {
	router := initialize.InitServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReadUserFailed(t *testing.T) {
	router := initialize.InitServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/0", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
