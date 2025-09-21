package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/auth0"
)

const (
	key    = "secret"
	MaxAge = 86400 * 30
	IsProd = false
)

func Auth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("AUTH0_KEY")
	clientSecret := os.Getenv("AUTH0_SECRET")
	domain := os.Getenv("AUTH0_DOMAIN")

	store := sessions.NewCookieStore([]byte(key))
	store.Options.MaxAge = MaxAge
	store.Options.Secure = IsProd
	store.Options.HttpOnly = true
	store.Options.SameSite = http.SameSiteLaxMode

	gothic.Store = store

	goth.UseProviders(
		auth0.New(
			clientID,
			clientSecret,
			"http://localhost:3000/auth/auth0/callback",
			domain,
		),
	)
}
