package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

// ReadAllDishes   GET "/api/v1/posts/readall"
func ReadAllDishes(ctx *gin.Context) {
	var dish []models.Dish
	var db gorm.DB = *(models.GetDB())
	db.Table("Dishes").Find(&dish)
	fmt.Println(dish)
	ctx.JSON(http.StatusOK, dish)
}
