package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jigintern/Foodmates-server/models"
	"github.com/jigintern/Foodmates-server/test/initialize"
	"github.com/stretchr/testify/assert"
)

func TestGetPostsSucceed(t *testing.T) {
	router := initialize.InitServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/posts/readall/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreatePostSucceed(t *testing.T) {
	jsonStr := `{"user_id:"` + "2" + `","dish_id":"` + "6" + `","comment":"` + "fafafafafafafafafa" + `","image_address":"` + "hahaha.png" + `"}`

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/posts/create/", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatalf("generate request failed.")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("send request failed.   %s", err.Error())
	}
	defer resp.Body.Close()

	resbin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read responce failed.")
	}
	var resjson models.Post
	json.Unmarshal(resbin, resjson)
	fmt.Println(resjson)
}

func TestCreatePostFailed(t *testing.T) {
	router := initialize.InitServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/posts/create/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
