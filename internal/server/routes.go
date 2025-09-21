package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)
	r.Get("/auth/{provider}/callback", s.getAuthCallBackHandler)
	r.Get("/auth/{provider}", s.getLoginHandler)
	r.Get("/logout/{provider}", s.getLogoutHandler)
	return r
}

var indexTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Auth0 SSO Demo</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
        .login-btn { 
            display: inline-block; 
            padding: 10px 20px; 
            background-color: #0066cc; 
            color: white; 
            text-decoration: none; 
            border-radius: 5px; 
        }
        .login-btn:hover { background-color: #0052a3; }
    </style>
</head>
<body>
    <h1>Auth0 SSO Demo</h1>
    <p>Click the button below to authenticate with Auth0:</p>
    <a href="/auth/auth0" class="login-btn">Login with Auth0</a>
</body>
</html>
`

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("index").Parse(indexTemplate)
	t.Execute(w, nil)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

var userTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>User Profile - Auth0 SSO</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
        .profile-info { background-color: #f5f5f5; padding: 15px; border-radius: 5px; margin: 10px 0; }
        .logout-btn { 
            display: inline-block; 
            padding: 8px 16px; 
            background-color: #dc3545; 
            color: white; 
            text-decoration: none; 
            border-radius: 3px; 
            margin-bottom: 20px;
        }
        .logout-btn:hover { background-color: #c82333; }
        .avatar { border-radius: 50%; max-width: 100px; }
    </style>
</head>
<body>
    <h1>Welcome, {{.Name}}!</h1>
    
    <a href="/logout/auth0" class="logout-btn">Logout</a>
    
    <div class="profile-info">
        <h2>Profile Information</h2>
        <p><strong>Name:</strong> {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
        <p><strong>Email:</strong> {{.Email}}</p>
        <p><strong>Username:</strong> {{.NickName}}</p>
        <p><strong>Location:</strong> {{.Location}}</p>
        <p><strong>Description:</strong> {{.Description}}</p>
        <p><strong>User ID:</strong> {{.UserID}}</p>
        {{if .AvatarURL}}
        <p><strong>Avatar:</strong><br>
           <img src="{{.AvatarURL}}" alt="User Avatar" class="avatar">
        </p>
        {{end}}
    </div>
    
    <div class="profile-info">
        <h2>Token Information</h2>
        <p><strong>Access Token:</strong> {{.AccessToken}}</p>
        <p><strong>Expires At:</strong> {{.ExpiresAt}}</p>
        <p><strong>Refresh Token:</strong> {{.RefreshToken}}</p>
    </div>
</body>
</html>
`

func (s *Server) getAuthCallBackHandler(w http.ResponseWriter, r *http.Request) {

	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)
}

func (s *Server) getLogoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	err := gothic.Logout(w, r)
	if err != nil {
		log.Printf("Logout error: %v", err)
	}

	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	returnTo := "http://localhost:3000"
	logoutURL := fmt.Sprintf("https://%s/v2/logout?client_id=%s&returnTo=%s",
		auth0Domain,
		url.QueryEscape(os.Getenv("AUTH0_KEY")),
		url.QueryEscape(returnTo),
	)

	w.Header().Set("Location", logoutURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (s *Server) getLoginHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}
