// models/user.go (OAuth Only version)
package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id             int64
	FirstName      string
	LastName       string
	FullName       string
	Username       string
	Email          string
	Organization   sql.NullString
	Provider       string
	ProviderUserID string
	AvatarURL      sql.NullString
	VerifiedEmail  bool
	CreatedAt      time.Time
}
