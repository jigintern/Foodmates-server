package main

import (
	"errors"
	"fmt"
	"github.com/jigintern/Foodmates-server/models"
	"github.com/jigintern/Foodmates-server/routers"
	"github.com/joho/godotenv"
	"net/http"
	"time"
)

func EnvLoad() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func main() {
	fmt.Printf("============= Started Foodmates-Server!! =============\n")

	err := EnvLoad()
	if err != nil {
		return
	}
	fmt.Printf("* .env loaded.\n")

	models.InitDB()
	fmt.Printf("* Models initialized.\n")

	router := routers.InitRouter()
	fmt.Printf("* Routers initialized.\n")

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
	fmt.Printf("* Now server is listening!!\n")

}
