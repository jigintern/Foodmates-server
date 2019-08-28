package initialize

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"github.com/jigintern/Foodmates-server/routers"
	"github.com/joho/godotenv"
	"log"
)

func EnvLoad() error {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return errors.New("code must be hoge")
	}
	return nil
}

func InitServer() *gin.Engine {
	EnvLoad()
	models.InitDB()
	return routers.InitRouter()
}
