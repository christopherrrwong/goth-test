package auth

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"

	"gothtest/internal/config"

	"github.com/gorilla/sessions"
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

func Auth(config *config.Config) error {

	var providers []goth.Provider

	auth0ClientID := config.Auth0.ClientID
	auth0ClientSecret := config.Auth0.ClientSecret
	auth0Domain := config.Auth0.Domain
	auth0CallbackURL := config.Auth0.CallbackURL

	validAuth0Config := auth0ClientID != "" && auth0ClientSecret != "" && auth0Domain != "" && auth0CallbackURL != ""

	if validAuth0Config {
		providers = append(providers, auth0.New(
			auth0ClientID,
			auth0ClientSecret,
			auth0CallbackURL,
			auth0Domain,
		))
	}

	var MaxAge = config.Session.MaxAge
	var IsProd = config.Session.IsProd
	sessionKey, err := GenerateRandomKey(32)
	if err != nil {
		log.Fatal("Error generating random key: %v", err)
	}

	var store = sessions.NewCookieStore([]byte(sessionKey))
	store.Options.MaxAge = MaxAge
	store.Options.Secure = IsProd
	store.Options.HttpOnly = config.Session.HttpOnly
	store.Options.SameSite = http.SameSiteLaxMode
	store.Options.Domain = "localhost" // Allow cookie to work across different ports on localhost
	gothic.Store = store

	goth.UseProviders(
		providers...,
	)

	return nil
}
