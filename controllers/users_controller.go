package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"log"
	"net/http"
)

// ReadSpecificUser   GET "/api/v1/users/:id/"
func ReadSpecificUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	loginName := ctx.Param("login_name")
	var userData models.User
	db, err := models.GetDB()

	// DBがなければ500を返す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	recordNotFound := db.Table("Users").Where("login_name = ?", loginName).First(&userData).RecordNotFound()
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

// SignUp   POST "/api/v1/users/signup"
func SignUp(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var signupParams models.SignUpParams
	err := ctx.BindJSON(&signupParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		log.Fatalln(err)
	}
	db, err := models.GetDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Fatalln(err)
	}
	tx := db.Table("Users").Begin()
	tx.Create(&signupParams)
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
	ctx.JSON(http.StatusOK, nil)
	fmt.Println("\x1b[32msuccess!!\x1b[0m")
	fmt.Println(signupParams)
}

// SignIn   POST "/api/v1/users/signin"
func SignIn(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var signInParams models.SignInParams
	var userData models.User
	err := ctx.BindJSON(&signInParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		log.Fatalln(err)
	}
	db, err := models.GetDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Fatalln(err)
	}
	recordNotFound := db.Table("Users").Where("login_name=?", signInParams.LoginName).First(&userData).RecordNotFound()
	if recordNotFound {
		ctx.JSON(http.StatusForbidden, nil)
		fmt.Println("Login incorrect")
		return
	} else {
		ctx.JSON(http.StatusOK, userData)
	}
}
