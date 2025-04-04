package handlers

import (
	"net/http"

	"github.com/arjkashyap/erlic.ai/internal/db/repositories"
	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/utils"
	"github.com/markbates/goth/gothic"
)

type ProfileHandler struct {
	UserRepository *repositories.UserRepository
}

func NewProfileHandler(ur *repositories.UserRepository) *ProfileHandler {
	return &ProfileHandler{
		UserRepository: ur,
	}
}

// GetProfile returns the user's profile data
func (ph *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := gothic.Store.Get(r, gothic.SessionName)

	// Get user data from repository
	user, err := ph.UserRepository.GetUserByEmail(session.Values["email"].(string))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to fetch user profile"}, nil, logger.Logger)
		return
	}

	// Return profile data
	profileData := utils.Envelope{
		"user_id":    user.Id,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"full_name":  user.FullName,
		"username":   user.Username,
		"avatar_url": user.AvatarURL.String,
		"provider":   user.Provider,
		"created_at": user.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusOK, profileData, nil, logger.Logger)
}
