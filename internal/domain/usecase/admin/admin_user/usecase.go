package admin_user

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type UseCaseAdminUser interface {
	GetUsers(ctx context.Context, sessionID, accessToken string) ([]entity.User, error)
	ResetPassword(ctx context.Context, sessionID, accessToken string, userID int64) (*string, error)
	UpdateRole(ctx context.Context, sessionID, accessToken, roleName string, userID int64) error
}

type useCaseAdminUser struct {
	ur       UserRepository
	validate Validate
	pass     PassHash
}

func NewUseCaseAdminUser(ur UserRepository, validate Validate, pass PassHash) UseCaseAdminUser {
	return &useCaseAdminUser{
		ur:       ur,
		validate: validate,
		pass:     pass,
	}
}
