package main

import (
	"fmt"
	"net/http"

	"./routers"
)

func main() {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        router,
		ReadTimeout:    60000,
		WriteTimeout:   60000,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
