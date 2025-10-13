package auth

import (
	"net/http"

	"sso-auth/internal/config"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/auth0"
	"github.com/markbates/goth/providers/azuread"
	"github.com/markbates/goth/providers/google"
)

func Auth(config *config.Config) error {

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

	//only auth0 is working at the moment for testing purposes.
	goth.UseProviders(
		auth0.New(
			config.Auth0.ClientID,
			config.Auth0.ClientSecret,
			config.Auth0.CallbackURL,
			config.Auth0.Domain,
		),
		google.New(
			config.Google.GoogleKey,
			config.Google.GoogleSecret,
			config.Google.CallbackURL,
		),
		azuread.New(
			config.AzureAD.AzureADKey,
			config.AzureAD.AzureADSecret,
			config.AzureAD.CallbackURL,
			nil,
		),
	)

	return nil
}
