package main

import (
	"fmt"
	"net/http"
	"time"

	"log"
	"errors"
	"./models"
	"./routers"
	"github.com/joho/godotenv"
)

func EnvLoad() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return errors.New("code must be hoge")
	}
	return nil
}

func main() {
	EnvLoad()
	models.InitDB()
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
