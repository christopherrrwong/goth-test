package main

import (
	"fmt"
	"log"
	"os"
	"sso-auth/internal/auth"
	"sso-auth/internal/config"
	"sso-auth/internal/database"
	"sso-auth/internal/server"
)

func main() {

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	config, err := config.LoadConfigForEnv(env)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.NewConnection(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	if err := auth.Auth(config); err != nil {
		log.Fatalf("Failed to initialize auth: %v", err)
	}

	srv := server.NewServer(config)

	fmt.Printf("Environment: %s\n", env)
	fmt.Printf("Server is running on port %d\n", config.Server.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
