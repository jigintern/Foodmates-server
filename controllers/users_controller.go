package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"net/http"
	"strconv"
)

// ReadSpecificUser   GET "/api/v1/users/:id/"
func ReadSpecificUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	if id == 0 {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	var userData models.User
	db, err := models.GetDB()
	
	// DBがなければ500を返す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	db.Table("Users").Where("user_id = ?", id).First(&userData)
	ctx.JSON(http.StatusOK, userData)
}

// ReadAllUsers   GET "/api/v1/users/"
func ReadAllUsers(ctx *gin.Context) {
	var userData []models.User
	db, err := models.GetDB()
	
	// DBがなければ500を返す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	db.Table("Users").Find(&userData)
	ctx.JSON(200, userData)
}
