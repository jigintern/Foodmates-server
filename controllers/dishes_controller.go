package controllers

import (
	"../models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// ReadAllDishes   GET "/api/v1/posts/readall"
func ReadAllDishes(ctx *gin.Context) {
	var dish []models.Dish
	var db gorm.DB = *(models.GetDB())
	fmt.Printf("db_addr____controller: %v\n", db)
	db.Table("Dishes").Find(&dish)
	fmt.Println(dish)
	ctx.JSON(http.StatusOK, dish)
}
