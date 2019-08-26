package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var dummyUserData = []gin.H{
	{
		"icon_address":        "public/img/users/icon/1.jpg",
		"user_name":           "watano",
		"following_user_size": 22,
		"favorite_dishes":     []int{2, 4, 12},
		"biography":           "ラーメン大好き！！",
		"posts_size":          44,
		"birthday":            time.Date(2000, 11, 18, 0, 0, 0, 0, time.UTC),
		"sex":                 0,
		"country":             "Japan",
		"prefecture":          "Kanagawa",
	},
	{
		"icon_address":        "public/img/users/icon/8.jpg",
		"user_name":           "でみ",
		"following_user_size": 71,
		"favorite_dishes":     []int{1, 2, 3, 4, 7, 8, 12},
		"biography":           "任意のオタク",
		"posts_size":          928,
		"birthday":            time.Date(1999, 1, 31, 0, 0, 0, 0, time.UTC),
		"sex":                 0,
		"country":             "Japan",
		"prefecture":          "Hokkaido",
	},
	{
		"icon_address":        "public/img/users/icon/19.jpg",
		"user_name":           "箒コウモリ",
		"following_user_size": 32,
		"favorite_dishes":     []int{12, 13, 14},
		"biography":           "でざいなー担当",
		"posts_size":          334,
		"birthday":            time.Date(2000, 3, 19, 0, 0, 0, 0, time.UTC),
		"sex":                 0,
		"country":             "Japan",
		"prefecture":          "Kagawa",
	},
	{
		"icon_address":        "public/img/users/icon/11.jpg",
		"user_name":           "はたはた",
		"following_user_size": 98,
		"favorite_dishes":     []int{},
		"biography":           "Vue完全に理解した",
		"posts_size":          0,
		"birthday":            time.Date(1820, 6, 2, 0, 0, 0, 0, time.UTC),
		"sex":                 0,
		"country":             "Japan",
		"prefecture":          "Tokyo",
	},
}

// ReadUsers   GET "/api/v1/users/:id/"
func ReadUsers(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	ctx.JSON(200, dummyUserData[id])
}
