package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	marshalled, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("impossible to marshall teacher: %s", err)
	}

	req, err := http.NewRequest("POST", "https://webhook.site/3e9d8b06-2925-4171-bc55-7e337e8be91e", bytes.NewReader(marshalled))
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("impossible to do request: %s", err)
	}

	fmt.Printf("Response: %+v\n", resp)

	// if user.Provider == "auth0" {
	// 	err = database.InsertSSOIntegrationMapping(user.Name, user.UserID, "")
	// 	if err != nil {
	// 		fmt.Fprintln(w, err)
	// 		return
	// 	}
	// }
}

func (s *Server) getLoginHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothic.BeginAuthHandler(w, r)
}
