package auth

import (
	"net/http"
	"os"

	"crypto/rand"
	"fmt"
	"log"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/auth0"
)

func GenerateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed to read random bytes: %w", err)
	}
	return key, nil
}

func Auth() {
	var providers []goth.Provider
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	auth0ClientID := os.Getenv("AUTH0_CLIENT_ID")
	auth0ClientSecret := os.Getenv("AUTH0_SECRET")
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	auth0CallbackURL := os.Getenv("AUTH0_CALLBACK_URL")

	validAuth0Config := auth0ClientID != "" && auth0ClientSecret != "" && auth0Domain != "" && auth0CallbackURL != ""

	if validAuth0Config {
		providers = append(providers, auth0.New(
			auth0ClientID,
			auth0ClientSecret,
			auth0CallbackURL,
			auth0Domain,
		))
	}

	var MaxAge = 86400 * 30
	var IsProd = false
	sessionKey, err := GenerateRandomKey(32)
	if err != nil {
		log.Fatal("Error generating random key: %v", err)
	}

	var store = sessions.NewCookieStore([]byte(sessionKey))
	store.Options.MaxAge = MaxAge
	store.Options.Secure = IsProd
	store.Options.HttpOnly = true
	store.Options.SameSite = http.SameSiteLaxMode
	gothic.Store = store

	goth.UseProviders(
		providers...,
	)
}
