package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"net/http"
	"strconv"
)

// ReadSpecificUser   GET "/api/v1/users/:id/"
func ReadSpecificUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	var userData models.User
	db, err := models.GetDB()

	// DBがなければ500を返す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	recordNotFound := db.Table("Users").Where("id = ?", id).First(&userData).RecordNotFound()
	if recordNotFound {
		ctx.JSON(http.StatusBadRequest, nil)
	} else {
		ctx.JSON(http.StatusOK, userData)
	}
}

// ReadAllUsers   GET "/api/v1/users/"
func ReadAllUsers(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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
