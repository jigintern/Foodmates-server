package initialize

import (
	"log"
	"errors"

	"../../routers"
	"../../models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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