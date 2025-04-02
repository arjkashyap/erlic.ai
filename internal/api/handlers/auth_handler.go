package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/arjkashyap/erlic.ai/internal/customerrors"
	"github.com/arjkashyap/erlic.ai/internal/db/repositories"
	"github.com/arjkashyap/erlic.ai/internal/env"
	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type AuthHandler struct {
	UserReposityory *repositories.UserRepository
}

func NewAuthHandler(ur *repositories.UserRepository) *AuthHandler {
	setupAuthProviders()

	return &AuthHandler{UserReposityory: ur}
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
	store.Options.Domain = ""

	gothic.Store = store

	goth.ClearProviders()
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, callbackUrl, "email", "profile"),
	)
}

func (ah *AuthHandler) AuthInitiate(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), gothic.ProviderParamKey, provider))

	gothic.BeginAuthHandler(w, r)
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

	new_user := utils.ConvertGothUserToAppUser(user)

	// Persist User
	existing_usr, err := ah.UserReposityory.GetUserByEmail(new_user.Email)
	if err != nil {
		err_msg := fmt.Sprintf("Error Fetching user from storage. \nError: %s", err)
		logger.Logger.Error(err_msg)
		customerrors.ErrorResponse(w, r, http.StatusInternalServerError, err_msg)
		return
	}
	if existing_usr == nil {
		if err := ah.UserReposityory.Create(new_user); err != nil {
			logger.Logger.Error("Unable to add user to Database error %s", err)
			customerrors.ErrorResponse(w, r, http.StatusInternalServerError, fmt.Sprintf("Error Adding User to Storage. \nError: %s", err))
			return
		}
		logger.Logger.Info("New user created", "email", new_user.Email)

	}

	// Session Management
	session, err := gothic.Store.Get(r, gothic.SessionName)
	if err != nil {
		logger.Logger.Error("Failed to get session", "error", err)
		http.Error(w, fmt.Sprintf("Session error: %v", err), http.StatusInternalServerError)
		return
	}
	// Store essential info in the session
	session.Values["username"] = new_user.Username
	session.Values["email"] = new_user.Email

	err = session.Save(r, w)
	if err != nil {
		logger.Logger.Error("Failed to save session", "error", err)
		http.Error(w, fmt.Sprintf("Session save error: %v", err), http.StatusInternalServerError)
		return
	}

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
