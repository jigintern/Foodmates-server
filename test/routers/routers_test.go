package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jigintern/Foodmates-server/test/initialize"
	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {
	router := initialize.InitServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/posts/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
