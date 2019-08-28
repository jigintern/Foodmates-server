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
	db := *models.GetDB()
	db.Table("Users").Where("user_id = ?", id).First(&userData)
	ctx.JSON(http.StatusOK, userData)
}

// ReadAllUsers   GET "/api/v1/users/"
func ReadAllUsers(ctx *gin.Context) {
	var userData []models.User
	db := *models.GetDB()
	db.Table("Users").Find(&userData)
	ctx.JSON(200, userData)
}
