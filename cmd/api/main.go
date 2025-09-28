package main

import (
	"gothtest/internal/auth"
	"gothtest/internal/server"
	"log"
	"net/http"
)

func main() {
	auth.Auth()

	server := &http.Server{
		Addr:    ":3000",
		Handler: server.NewServer().Handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
