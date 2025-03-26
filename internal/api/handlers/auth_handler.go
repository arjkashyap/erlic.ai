package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/arjkashyap/erlic.ai/internal/env"
	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	setupAuthProviders()

	return &AuthHandler{}
}

func setupAuthProviders() {
	logger.Logger.Info("Registering Provider with Goth")
	googleClientId := env.GetString("GOOGLE_CLIENT_ID", "")
	googleClientSecret := env.GetString("GOOGLE_CLIENT_SECRET", "")

	callbackUrl := "http://localhost:8080/auth/google/callback"

	key := env.GetString("SESSION_SECRET", "secretkey")
	isProd := env.GetBool("IS_PRODUCTION", false)

	if googleClientId == "" || googleClientSecret == "" {
		panic("Google credentials invalid or empty")
	}

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(86400 * 30) // 30 days
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd
	store.Options.SameSite = http.SameSiteLaxMode

	gothic.Store = store

	goth.ClearProviders()
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, callbackUrl, "email", "profile"),
	)
}

func (ah *AuthHandler) AuthInitiate(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), gothic.ProviderParamKey, provider))

	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		logger.Logger.Info(gothUser)
		http.Redirect(w, r, "http://localhost:5173/dashboard", http.StatusFound)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (ah *AuthHandler) AuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), gothic.ProviderParamKey, provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		logger.Logger.Error("Auth callback error", "error", err)
		http.Error(w, fmt.Sprintf("Authentication error: %v", err), http.StatusInternalServerError)
		return
	}

	logger.Logger.Info("User authenticated", "user", user)

	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}

// AuthLogout handles user logout
func (ah *AuthHandler) AuthLogout(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), gothic.ProviderParamKey, provider))

	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
