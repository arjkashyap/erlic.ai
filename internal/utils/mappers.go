package utils // Or another appropriate package

import (
	"database/sql"
	"strings"

	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/models"
	"github.com/markbates/goth"
)

func ConvertGothUserToAppUser(gothUser goth.User) *models.User {
	// Generate username from email
	emailParts := strings.Split(gothUser.Email, "@")
	username := emailParts[0]

	// Split full name into first and last name
	nameParts := strings.Fields(gothUser.Name)
	firstName := ""
	lastName := ""

	if len(nameParts) > 0 {
		firstName = nameParts[0]
		if len(nameParts) > 1 {
			lastName = strings.Join(nameParts[1:], " ")
		}
	}

	avatarURL := sql.NullString{Valid: false}
	if gothUser.AvatarURL != "" {
		avatarURL = sql.NullString{String: gothUser.AvatarURL, Valid: true}
	}

	appUser := &models.User{
		FirstName:      firstName,
		LastName:       lastName,
		FullName:       gothUser.Name,
		Username:       username,
		Email:          gothUser.Email,
		Organization:   sql.NullString{Valid: false},
		Provider:       gothUser.Provider,
		ProviderUserID: gothUser.UserID,
		AvatarURL:      avatarURL,
		VerifiedEmail:  true,
	}

	logger.Logger.Info("user mapped", "email", appUser.Email, "username", appUser.Username)
	return appUser
}
