package auth

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	UpdateUser(id int64, user *entity.UpdateUser) error
	UpdateUserRole(id int64, role string) error
	DeleteUser(id int64) error
	GetUserById(id int64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type SessionRepository interface {
	Create(ctx context.Context, session entity.Session) (string, error)
	Get(ctx context.Context, sessionID string) (*entity.Session, error)
	Update(ctx context.Context, sessionID string, session entity.Session) error
	Delete(ctx context.Context, sessionID string) error
}

type Jwt interface {
	GenerateToken(payload entity.Payload) (string, string, error)
	ValidateJwt(token string) (*entity.Payload, error)
}

type PassHash interface {
	EncryptPassword(password string) (string, error)
	CheckPass(password, dbPassword string) error
}
