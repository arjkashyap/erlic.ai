package repositories

import (
	"database/sql"

	"github.com/arjkashyap/erlic.ai/internal/models"
)

type UserRepositoryInterface interface {
	Create(*models.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
