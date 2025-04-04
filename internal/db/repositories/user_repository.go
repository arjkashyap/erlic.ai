package repositories

import (
	"database/sql"
	"fmt"

	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/arjkashyap/erlic.ai/internal/models"
)

type UserRepositoryInterface interface {
	Create(*models.User) error
	GetUserByEmail(string) (*models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	logger.Logger.Info("Adding user to storage")
	query := `
    INSERT INTO users (
        first_name, last_name, full_name, username, 
        email, organization, provider, provider_user_id, 
        avatar_url, verified_email
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
    ) RETURNING id, created_at`

	err := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.FullName,
		user.Username,
		user.Email,
		user.Organization,
		user.Provider,
		user.ProviderUserID,
		user.AvatarURL,
		user.VerifiedEmail,
	).Scan(&user.Id, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	logger.Logger.Info("added user to store",
		"username", user.Username,
		"id", user.Id)
	return user, nil
}

func (ur *UserRepository) GetUserByID(id int64) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, first_name, last_name, username, email, organization,
               provider, provider_user_id, avatar_url, verified_email, created_at
        FROM users
        WHERE id = $1`

	err := ur.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Organization,
		&user.Provider,
		&user.ProviderUserID,
		&user.AvatarURL,
		&user.VerifiedEmail,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Logger.Warn("User not found by ID", "id", id)
			return nil, sql.ErrNoRows // Indicate not found clearly
		}
		logger.Logger.Error("Error fetching user by ID", "id", id, "error", err)
		return nil, fmt.Errorf("error fetching user by ID %d: %w", id, err)
	}

	// Construct FullName after scanning FirstName and LastName
	user.FullName = user.FirstName + " " + user.LastName

	logger.Logger.Debug("Found user by ID", "id", user.Id)
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := `SELECT id, first_name, last_name, username, email, organization, provider, provider_user_id, avatar_url, verified_email 
              FROM users 
              WHERE email = $1`

	var org sql.NullString
	var avatarURL sql.NullString

	err := ur.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Organization,
		&user.Provider,
		&user.ProviderUserID,
		&avatarURL,
		&user.VerifiedEmail,
	)
	user.FullName = user.FirstName + user.LastName

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	user.Organization = org
	user.AvatarURL = avatarURL
	logger.Logger.Infof("Found user %s in the db", user.Email)
	return user, nil
}
