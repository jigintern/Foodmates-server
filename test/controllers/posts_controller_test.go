package controllers_test

import (
	"encoding/json"
	"fmt"
	"github.com/jigintern/Foodmates-server/test/initialize"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"strconv"
	"bytes"

	"github.com/jigintern/Foodmates-server/models"
	"github.com/stretchr/testify/assert"
)

// TestReadAllPosts   全ての投稿を取得するAPIのテスト
func TestReadAllPosts(t *testing.T) {
	
	// 初期処理
	
	// 成功ケース
	t.Run("succeed", func(t *testing.T) {
		
		// 初期処理
		request, err := http.NewRequest("GET", "http://localhost:8080/api/v1/posts/readall/", nil)
		if err != nil { t.Fatalf("\x1b[31msend request failed. (%s)\x1b[0m\n", err.Error()) }
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil { t.Fatalf("\x1b[31mrecieve request failed. (%s)\x1b[0m\n", err.Error()) }

		defer resp.Body.Close()
		resbin, err := ioutil.ReadAll(resp.Body)
		var resjson []models.Post
		err = json.Unmarshal(resbin, &resjson)
		if err != nil { t.Fatalf("\x1b[31mjson unmarshal failed. (%s)\x1b[0m\n", err.Error()) }
		t.Logf("\n\x1b[33m======= responce =======\n%+v\x1b[0m\n", resjson)
		// ステータスコード200が返されること
		t.Run("return response code 200", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})

		// 全ての投稿データがPostの配列として返されること
		t.Run("response type is []Post", func(t *testing.T) {
			t.Parallel()
			assert.NotEmpty(t, resjson)
		})
		// 配列の全要素の形式が正しいこと
		t.Run("all of response datas are correct format", func(t *testing.T) {
			t.Parallel()
			var resjson models.Post
			err = json.Unmarshal(resbin, &resjson)
			fmt.Println(resjson)
		})
	})

	// 終了処理

}

// TestReadSpecificUsersPost   特定のユーザーの投稿のみを全て取得するAPIのテスト
func TestReadSpecificUsersPost(t *testing.T) {
	
	// 初期処理
	
	// 成功ケース
	t.Run("succeed", func(t *testing.T) {
		
		// 初期処理
		testUserId := 1
		request, err := http.NewRequest("GET", "http://localhost:8080/api/v1/posts/read/" + strconv.Itoa(testUserId), nil)
		if err != nil { t.Fatalf("\x1b[31msend request failed. (%s)\x1b[0m\n", err.Error()) }
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil { t.Fatalf("\x1b[31mrecieve request failed. (%s)\x1b[0m\n", err.Error()) }

		defer resp.Body.Close()
		resbin, err := ioutil.ReadAll(resp.Body)
		var resjson []models.Post
		err = json.Unmarshal(resbin, &resjson)
		if err != nil { t.Fatalf("\x1b[31mjson unmarshal failed. (%s)\x1b[0m\n", err.Error()) }
		t.Logf("\n\x1b[33m======= responce =======\n%+v\x1b[0m\n", resjson)

		// ステータスコード200が返されること
		t.Run("return response code 200", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})

		// 全ての投稿データがPostの配列として返されること
		t.Run("response type is []Post", func(t *testing.T) {
			t.Parallel()
			assert.NotEmpty(t, resjson)
		})

		// 配列の全要素の形式が正しいこと
		t.Run("all of response datas are correct format", func(t *testing.T) {
			t.Parallel()
		})

		// 全ての投稿データのuser_idが指定したIDであること
		t.Run("all of response data's user_id is expected value", func(t *testing.T) {
			t.Parallel()
			for i := 0; i < len(resjson); i++ {
				assert.Equal(t, testUserId, resjson[i].UserId)
			}
		})
	})

	// 終了処理

}

// TestReadAllPosts   投稿を追加するAPIのテスト
func TestCreatePost(t *testing.T) {
	
	// 初期処理
	
	// 成功ケース
	t.Run("succeed", func(t *testing.T) {
		
		// 初期処理
		testdata := models.Post{
			UserId: 2,
			DishId: 4,
			Comment: "it was delicious!",
			ImageAddress: "img/gazou/abc/123.jpg",
		}
		testdata_bin, err := json.Marshal(testdata)
		if err != nil { t.Fatalf("\x1b[31mjson marshal failed. (%s)\x1b[0m\n", err.Error()) }
		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/posts/create/", bytes.NewBuffer([]byte(testdata_bin)))
		if err != nil { t.Fatalf("\x1b[31msend request failed. (%s)\x1b[0m\n", err.Error()) }
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil { t.Fatalf("\x1b[31mrecieve request failed. (%s)\x1b[0m\n", err.Error()) }
		
		defer resp.Body.Close()
		resbin, err := ioutil.ReadAll(resp.Body)
		var resjson models.Post
		err = json.Unmarshal(resbin, &resjson)
		t.Logf("\n\x1b[33m======= responce =======\n%+v\x1b[0m\n", resjson)
		
		// ステータスコード200が返されること
		t.Run("return response code 200", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})

		// 送信したデータが返されること
		t.Run("return response array of all posts", func(t *testing.T) {
			t.Parallel()
			sendedData := models.Post{
				UserId: resjson.UserId,
				DishId: resjson.DishId,
				Comment: resjson.Comment,
				ImageAddress: resjson.ImageAddress,
			}
			assert.Equal(t, testdata, sendedData)
		})
		
		// ReadAllPostsで、作成した内容のデータが末尾に追加されて返ってくること
		t.Run("when read all posts, created data is returned at last of array", func(t *testing.T) {
			req_t, err := http.NewRequest("GET", "http://localhost:8080/api/v1/posts/readall/", nil)
			if err != nil { t.Fatalf("\x1b[31msend request failed. (%s)\x1b[0m\n", err.Error()) }
			client_t := &http.Client{}
			resp_t, err := client_t.Do(req_t)
			if err != nil { t.Fatalf("\x1b[31mrecieve request failed. (%s)\x1b[0m\n", err.Error()) }
			
			defer resp_t.Body.Close()
			resbin_t, err := ioutil.ReadAll(resp_t.Body)
			var resjson_t []models.Post
			err = json.Unmarshal(resbin_t, &resjson_t)
			if err != nil { t.Fatalf("\x1b[31mjson unmarshal failed. (%s)\x1b[0m\n", err.Error()) }
			t.Logf("\n\x1b[32m======= /api/v1/posts/readall/ response =======\n%+v\x1b[0m\n", resjson_t)
			
			lastData := models.Post{
				UserId: resjson_t[len(resjson_t)-1].UserId,
				DishId: resjson_t[len(resjson_t)-1].DishId,
				Comment: resjson_t[len(resjson_t)-1].Comment,
				ImageAddress: resjson_t[len(resjson_t)-1].ImageAddress,
			}
			assert.Equal(t, testdata, lastData)
		})
	})

	// 終了処理
}
