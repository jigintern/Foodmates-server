package controllers_test

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jigintern/Foodmates-server/models"
	"github.com/jigintern/Foodmates-server/test/initialize"
	"github.com/stretchr/testify/assert"
)

// TestReadAllPosts   全ての投稿を取得するAPIのテスト
func TestReadAllPosts(t *testing.T) {
	
	// 初期処理
	t.Logf("[TestReadAllPosts] setup: %s\n", time.Now())
	
	// 成功ケース
	t.Run("succeed", func(t *testing.T) {
		
		// 初期処理
		t.Logf("[TestReadAllPosts/success] setup: %s\n", time.Now())
		var (
			router *gin.Engine
			recorder *httptest.ResponseRecorder
			request *http.Request
			err error
		)
		//initialize.EnvLoad()
		//models.InitDB()
		//router = routers.InitRouter()
		recorder = httptest.NewRecorder()
		request, err = http.NewRequest("GET", "http://localhost:8080/api/v1/posts/readall/", nil)
		if err != nil {
			t.Fatalf("send request failed. (%s)\n", err.Error())
		}
		fmt.Printf("recorder: %v, request: %v\n", recorder, request)
		router.ServeHTTP(recorder, request)
		
		// ステータスコード200が返されること
		t.Run("return response code 200", func(t *testing.T) {
			//t.Parallel()
			assert.Equal(t, http.StatusOK, recorder.Code)
			t.Logf("[TestReadAllPosts] return response code 200: %s\n", time.Now())
		})

		// 全ての投稿データが配列として返されること
		t.Run("return response array of all posts", func(t *testing.T) {
			//t.Parallel()
			var resjson models.Post
			err = json.Unmarshal(recorder.Body.Bytes(), resjson)
			if err != nil {
				t.Fatalf("responce json unmarshal failed. (%s)\n", err.Error())
			}
			fmt.Println(resjson)
		})
	})

	// 失敗ケース
	t.Run("failed", func(t *testing.T) {
		
		// 初期処理
		t.Logf("[TestReadAllPosts/failed] setup: %s\n", time.Now())
		
		// データベースが見つからないときに、ステータスコード500が返されること
		t.Run("return response code 500 when database missing", func(t *testing.T) {
			//t.Parallel()
			var (
				router *gin.Engine
				recorder *httptest.ResponseRecorder
				request *http.Request
				err error
			)
			//initialize.EnvLoad()
			//router = routers.InitRouter()
			recorder = httptest.NewRecorder()
			request, err = http.NewRequest("GET", "http://localhost:8080/api/v1/posts/readall/", nil)
			if err != nil {
				t.Fatalf("send request failed. (%s)\n", err.Error())
			}
			router.ServeHTTP(recorder, request)
			
			assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			t.Logf("[TestReadAllPosts] return response code 500 when database missing: %s\n", time.Now())
		})
	})
	
	// 終了処理
	t.Logf("[TestReadAllPosts] tear-down: %s\n", time.Now())

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
