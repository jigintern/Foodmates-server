package main

import (
	"fmt"
	"net/http"
	"time"

	"log"
	"errors"
//	"./models"
	"github.com/jigintern/Foodmates-server/models"
	"github.com/jigintern/Foodmates-server/routers"
	//	"./routers"
	"github.com/joho/godotenv"
)

func EnvLoad() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
		return errors.New("Error loading .env file.")
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
