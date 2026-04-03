package admin_user

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	UpdateUserRole(id int64, role string) error
}

type SessionRepository interface {
	Get(ctx context.Context, sessionID string) (*entity.Session, error)
}

type Jwt interface {
	ValidateJwt(token string) (*entity.Payload, error)
}
