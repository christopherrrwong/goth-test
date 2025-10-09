package auth

import (
	"net/http"

	"sso-auth/internal/config"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/auth0"
)

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

	hashKey := securecookie.GenerateRandomKey(64)
	blockKey := securecookie.GenerateRandomKey(32)

	var store = sessions.NewCookieStore(hashKey, blockKey)
	store.Options.MaxAge = MaxAge
	store.Options.Secure = IsProd
	store.Options.HttpOnly = true
	store.Options.SameSite = http.SameSiteLaxMode
	gothic.Store = store

	goth.UseProviders(
		providers...,
	)

	return nil
}
