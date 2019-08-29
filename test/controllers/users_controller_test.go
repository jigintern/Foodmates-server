package controllers

import (
	"encoding/json"
	"github.com/jigintern/Foodmates-server/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUserSucceed(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		testUserId := 1
		request, err := http.NewRequest("GET", "http://localhost:8080/api/v1/users/"+strconv.Itoa(testUserId), nil)
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
		var resjson models.User
		err = json.Unmarshal(resbin, &resjson)
		if err != nil {
			t.Fatalf("\x1b[31mjson unmarshal failed. (%s)\x1b[0m\n", err.Error())
		}
		t.Logf("\n\x1b[33m======= responce =======\n%+v\x1b[0m\n", resjson)

		// ステータスコード200が返されること
		t.Run("return response code 200", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})

		// 全ての投稿データがPostの配列として返されること
		t.Run("response type is []User", func(t *testing.T) {
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
			assert.Equal(t, testUserId, resjson.ID)
		})
	})
}
