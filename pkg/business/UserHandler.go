package business

import (
	"rastebazaar/pkg/domain"
)

// UserHandler will hold everything that controller needs
type UserHandler struct {
	userRepo domain.UserRepository
}

// NewUserHandler returns a new BaseHandler
func NewUserHandler(userRepo domain.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}
