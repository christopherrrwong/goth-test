package server

import (
	"context"
	"net/http"
	"sso-auth/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   s.config.Cors.AllowedOrigins,
		AllowedMethods:   s.config.Cors.AllowedMethods,
		AllowedHeaders:   s.config.Cors.AllowedHeaders,
		AllowCredentials: s.config.Cors.AllowCredentials,
		MaxAge:           s.config.Cors.MaxAge,
	}))

	r.Get("/sso-auth/{provider}/callback", s.getAuthCallBackHandler)
	r.Get("/sso-auth/{provider}", s.getLoginHandler)

	return r
}

func (s *Server) getAuthCallBackHandler(w http.ResponseWriter, r *http.Request) {

	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	uuid, err := r.Cookie("uuid")

	var uuidValue string

	if uuid == nil {
		uuidValue = ""
	} else {
		uuidValue = uuid.Value
	}

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, "Error completing user auth", http.StatusUnauthorized)
		return
	}

	if user.Provider == "auth0" {
		err = database.InsertSSOIntegrationMapping(user.Name, uuidValue)
		if err != nil {
			http.Error(w, "Error inserting sso integration mapping", http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getLoginHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	if provider == "" {
		http.Error(w, "Error getting provider", http.StatusBadRequest)
		return
	}

	uuid := r.URL.Query().Get("uuid")

	http.SetCookie(w, &http.Cookie{
		Name:     "uuid",
		Value:    uuid,
		Path:     "/",
		MaxAge:   300,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Secure:   s.config.Session.IsProd,
	})

	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothic.BeginAuthHandler(w, r)

}
