package controllers

/*
import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jigintern/Foodmates-server/test/initialize"
	"github.com/stretchr/testify/assert"
)

func TestReadUserSucceed(t *testing.T) {
	router := initialize.InitServer()
	expectedResponseJSON := `{"id": 1,"created_at": "2019-08-27T07:36:43+09:00","updated_at": "2019-08-27T07:46:18+09:00","name": "hogehoge","biography": "","country": "Jap","prefecture": "Fukui","icon_address": ""}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedResponseJSON, w.Body.String())
}

func TestReadUserFailed(t *testing.T) {
	router := initialize.InitServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/0", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
*/