package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"net/http"
)

// ReadAllDishes   GET "/api/v1/dishes/readall"
func ReadAllDishes(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var dish []models.Dish
	db, err := models.GetDB()
	
	// DBがなければ500を返す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	db.Table("Dishes").Find(&dish)
	fmt.Println(dish)
	ctx.JSON(http.StatusOK, dish)
}

// CreateDish   GET "/api/v1/dishes/create"
func CreateDish(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var param models.Dish
	err := ctx.BindJSON(&param)
	if err != nil {
		fmt.Println("======== request couldn't bind to json!! ========")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "The request couldn't bind to json."})
		return
	}
	
	db, err := models.GetDB()
	tx := db.Table("Dishes").Begin()
	tx.Create(&param)
	if tx.Error != nil {
		fmt.Println("\x1b[31mstarting transition failed.\x1b[0m")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "starting transition failed."})
		return
	}
	
	if len(db.GetErrors()) != 0 {
		fmt.Println("\x1b[31mchanging database failed.\x1b[0m")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "changing database failed."})
		return
	}
	
	tx.Commit()
	ctx.JSON(http.StatusOK, param)
	fmt.Println("\x1b[32msuccess!!\x1b[0m")
	fmt.Println(param)
}
