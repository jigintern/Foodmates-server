package controllers

import (
	"github.com/jigintern/Foodmates-server/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ReadUsers   GET "/api/v1/users/:id/"
func ReadUsers(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	var userData models.User
	db := *models.GetDB()
	db.Table("Users").Where("user_id = ?", id).First(&userData)
	ctx.JSON(200, userData)
}
