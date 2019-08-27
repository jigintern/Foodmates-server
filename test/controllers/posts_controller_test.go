package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"

	"../initialize"
	"../../models"
	"github.com/stretchr/testify/assert"
)

func TestGetPostsSucceed(t *testing.T) {
	router := initialize.InitServer()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/posts/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreatePostSucceed(t *testing.T) {
	jsonStr := `{"user_id:"` + "2" + `","dish_id":"` + "6" + `","comment":"` + "fafafafafafafafafa" + `","image_address":"` + "hahaha.png" + `"}`

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/posts/", bytes.NewBuffer([]byte(jsonStr)))
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
	req, _ := http.NewRequest("POST", "/api/v1/posts/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}