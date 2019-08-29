package controllers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/jigintern/Foodmates-server/models"
	"github.com/stretchr/testify/assert"
)

// TestReadAllPosts   全ての投稿を取得するAPIのテスト
func TestReadAllPosts(t *testing.T) {
	
	// 初期処理
	t.Logf("\x1b[36m[TestReadAllPosts] setup: %s\x1b[0m\n", time.Now())
	
	// 成功ケース
	t.Run("succeed", func(t *testing.T) {
		
		// 初期処理
		t.Logf("\x1b[36m[TestReadAllPosts/success] setup: %s\x1b[0m\n", time.Now())
		
		request, err := http.NewRequest("GET", "http://localhost:8080/api/v1/posts/readall/", nil)
		if err != nil {
			t.Fatalf("\x1b[31msend request failed. (%s)\x1b[0m\n", err.Error())
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			t.Fatalf("\x1b[31mrecieve request failed. (%s)\x1b[0m\n", err.Error())
		}
		defer resp.Body.Close()
		resbin, err := ioutil.ReadAll(resp.Body)
		
		// ステータスコード200が返されること
		t.Run("return response code 200", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			t.Logf("\x1b[35m[TestReadAllPosts] return response code 200: %s\x1b[0m\n", time.Now())
		})

		// 全ての投稿データが配列として返されること
		t.Run("return response array of all posts", func(t *testing.T) {
			t.Parallel()
			var resjson models.Post
			err = json.Unmarshal(resbin, resjson)
			fmt.Println(resjson)
			t.Logf("\x1b[35m[TestReadAllPosts] return response array of all posts: %s\x1b[0m\n", time.Now())
		})
	})

	// 終了処理
	t.Logf("\x1b[36m[TestReadAllPosts] tear-down: %s\x1b[0m\n", time.Now())

}

/*
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
*/
