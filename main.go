package main

import (
	"fmt"
	"net/http"
	"time"

	"./routers"
)

func main() {
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
