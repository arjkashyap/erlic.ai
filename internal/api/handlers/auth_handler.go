package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/arjkashyap/erlic.ai/internal/customerrors"
	"github.com/arjkashyap/erlic.ai/internal/db/repositories"
	"github.com/arjkashyap/erlic.ai/internal/env"
	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/models"
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

	callbackUrl := "http://localhost:8080/api/auth/google/callback"

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

	logger.Logger.Info("Session Options", "options", store.Options) // Add this log

	goth.ClearProviders()
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, callbackUrl, "email", "profile"),
	)
}

// Enhanced GetCurrentUser checks session and returns full user details
func (ah *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	session, err := gothic.Store.Get(r, gothic.SessionName)

	if err != nil || session.IsNew || len(session.Values) == 0 {
		logger.Logger.Debug("GetCurrentUser: No valid session found error", err, " isNew", session != nil && session.IsNew)
		utils.WriteJSON(w, http.StatusOK, utils.Envelope{"authenticated": false, "user": nil}, nil, logger.Logger)
		return
	}

	userIDRaw, ok := session.Values["user_id"]
	if !ok {
		logger.Logger.Warn("GetCurrentUser: user_id not found in session values")
		utils.WriteJSON(w, http.StatusOK, utils.Envelope{"authenticated": false, "user": nil}, nil, logger.Logger)
		return
	}

	userID, ok := userIDRaw.(int64)
	if !ok {
		logger.Logger.Error("GetCurrentUser: user_id in session is not of expected type (int)", "value", userIDRaw)
		// Maybe clear the invalid session value?
		utils.WriteJSON(w, http.StatusOK, utils.Envelope{"authenticated": false, "user": nil}, nil, logger.Logger)
		return
	}

	user, err := ah.UserReposityory.GetUserByID(userID)
	if err != nil {
		// Handle cases where user ID from session doesn't exist in DB anymore
		if err == sql.ErrNoRows {
			logger.Logger.Warn("GetCurrentUser: User ID found in session but not in DB", "user_id", userID)
			// Consider clearing the invalid session here as well
			// gothic.Logout(w, r) // Or manually clear session.Values["user_id"] and Save
			utils.WriteJSON(w, http.StatusOK, utils.Envelope{"authenticated": false, "user": nil}, nil, logger.Logger)
		} else {
			// Handle other DB errors
			logger.Logger.Error("GetCurrentUser: Failed to fetch user by ID", "user_id", userID, "error", err)
			customerrors.ErrorResponse(w, r, http.StatusInternalServerError, "Failed to retrieve user data")
		}
		return
	}

	// --- User found and authenticated ---
	logger.Logger.Debug("GetCurrentUser: User is authenticated", "user_id", userID, "email", user.Email)
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"authenticated": true, "user": user}, nil, logger.Logger)
}

func (ah *AuthHandler) AuthInitiate(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), gothic.ProviderParamKey, provider))

	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		logger.Logger.Info(gothUser)
		http.Redirect(w, r, "http://localhost:3000/dashboard", http.StatusFound)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (ah *AuthHandler) AuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), gothic.ProviderParamKey, provider))

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		logger.Logger.Error("Auth callback error", "error", err)
		http.Error(w, fmt.Sprintf("Authentication error: %v", err), http.StatusInternalServerError)
		return
	}

	appUser := utils.ConvertGothUserToAppUser(gothUser)

	// persist User
	existingUser, err := ah.UserReposityory.GetUserByEmail(appUser.Email)
	if err != nil {
		errMsg := fmt.Sprintf("Error fetching user from storage: %s", err)
		logger.Logger.Error(errMsg)
		customerrors.ErrorResponse(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	var user *models.User
	if existingUser == nil {
		// Create new user if they don't exist
		user, err = ah.UserReposityory.Create(appUser)
		if err != nil {
			errMsg := fmt.Sprintf("Error adding user to storage: %s", err)
			logger.Logger.Error(errMsg)
			customerrors.ErrorResponse(w, r, http.StatusInternalServerError, errMsg)
			return
		}
		logger.Logger.Info("New user created", "email", user.Email)
	} else {
		user = existingUser
	}

	// Session Management
	session, err := gothic.Store.Get(r, gothic.SessionName)
	if err != nil {
		logger.Logger.Error("Failed to get session", "error", err)
		http.Error(w, fmt.Sprintf("Session error: %v", err), http.StatusInternalServerError)
		return
	}

	// Clear any existing session values
	session.Values = make(map[interface{}]interface{})

	session.Values["user_id"] = user.Id

	isProd := env.GetBool("IS_PRODUCTION", false)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30, // 30 days
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteLaxMode,
	}

	if err := session.Save(r, w); err != nil {
		logger.Logger.Error("Failed to save session", "error", err)
		http.Error(w, fmt.Sprintf("Session save error: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:3000/dashboard", http.StatusFound)
}

// AuthLogout handles user logout
func (ah *AuthHandler) AuthLogout(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), gothic.ProviderParamKey, provider))

	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
