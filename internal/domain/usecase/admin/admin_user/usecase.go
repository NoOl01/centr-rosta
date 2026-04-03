package admin_user

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type UseCaseAdminUser interface {
	GetUsers(ctx context.Context, sessionID, accessToken string) ([]entity.User, error)
}

type useCaseAdminUser struct {
	ur      UserRepository
	session SessionRepository
	jwt     Jwt
}

func NewUseCaseAdminUser(ur UserRepository, session SessionRepository, jwt Jwt) UseCaseAdminUser {
	return &useCaseAdminUser{
		ur:      ur,
		session: session,
		jwt:     jwt,
	}
}
