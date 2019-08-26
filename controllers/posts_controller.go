package controllers

import (
	"fmt"
	"strconv"
	"time"
	"../models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PostData struct {
	UserID  int    `json:"user_id"`
	Content string `json:"comment"`
	DishID  int    `json:"dish_id"`
}


var dummyData = []gin.H{
	{
		"dish_name":             "塩ラーメン",
		"restaurant_name":       "らあめん極",
		"restaurant_place":      "東京都 新宿区",
		"user_name":             "watano",
		"user_icon_address":     "public/img/users/icon/1.jpg",
		"is_bookmark":           false,
		"content":               "さっぱりしていて美味しかった。",
		"content_image_address": "public/img/posts/210.jpg",
		"created_at":            time.Date(2019, 8, 22, 10, 43, 22, 0, time.UTC),
	},
	{
		"dish_name":             "特製やきそば",
		"restaurant_name":       "麺麺",
		"restaurant_place":      "北海道 札幌市",
		"user_name":             "でみ",
		"user_icon_address":     "public/img/users/icon/2.jpg",
		"is_bookmark":           true,
		"content":               "量が多くて満足。また行きたい。",
		"content_image_address": "public/img/posts/222.jpg",
		"created_at":            time.Date(2019, 8, 22, 13, 11, 01, 0, time.UTC),
	},
	{
		"dish_name":             "虹色ハンバーグ",
		"restaurant_name":       "謎の店",
		"restaurant_place":      "沖縄県 沖縄市",
		"user_name":             "はたはた",
		"user_icon_address":     "public/img/users/icon/6.jpg",
		"is_bookmark":           false,
		"content":               "宇宙の味がした。",
		"content_image_address": "public/img/posts/309.jpg",
		"created_at":            time.Date(1029, 3, 1, 22, 01, 56, 0, time.UTC),
	},
	{
		"dish_name":             "わんこCOMP",
		"restaurant_name":       "jig.jp",
		"restaurant_place":      "福井県 鯖江市",
		"user_name":             "箒コウモリ",
		"user_icon_address":     "public/img/users/icon/3.jpg",
		"is_bookmark":           true,
		"content":               "一生分の栄養を摂取したような気分になった",
		"content_image_address": "public/img/posts/111.jpg",
		"created_at":            time.Date(2019, 8, 21, 7, 13, 01, 0, time.UTC),
	},
}


// ReadPosts   GET "/api/v1/posts"
func ReadPosts(ctx *gin.Context) {
	var post []models.Post
	var db gorm.DB = *(models.GetDB())
	fmt.Printf("db_addr____controller: %v\n", db)
	db.Table("Posts").Find(&post)
	fmt.Println(post)
	ctx.JSON(200, post)
}

// CreatePost   POST "/api/v1/posts"
func CreatePost(ctx *gin.Context) {
	var data PostData
	err := ctx.BindJSON(&data)
	if err != nil {
		fmt.Println("======== request couldn't bind to json!! ========")
		return
	}
	dummyData = append(dummyData, gin.H{
		"dish_name":             "ラーメン" + strconv.Itoa(data.DishID),
		"restaurant_name":       "美味しい店" + strconv.Itoa(data.DishID),
		"restaurant_place":      "東京都 " + strconv.Itoa(data.DishID) + "番区",
		"user_name":             "ユーザー" + strconv.Itoa(data.DishID),
		"user_icon_address":     "public/img/users/icon/" + strconv.Itoa(data.UserID) + ".jpg",
		"is_bookmark":           false,
		"content":               data.Content,
		"content_image_address": "public/img/posts/" + strconv.Itoa(data.DishID) + ".jpg",
		"created_at":            time.Now(),
	})
	ctx.JSON(200, gin.H{"data": data})
	fmt.Println("======== success!! ========")
	fmt.Println(data)
}
