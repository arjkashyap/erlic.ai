package handlers

import (
	"net/http"

	"github.com/arjkashyap/erlic.ai/internal/db/repositories"
)

type UserHandler struct {
	UserRepository *repositories.UserRepository
}

func NewUserHandler(ur *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: ur,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}
