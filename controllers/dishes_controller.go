package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"net/http"
)

// ReadAllDishes   GET "/api/v1/posts/readall"
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
