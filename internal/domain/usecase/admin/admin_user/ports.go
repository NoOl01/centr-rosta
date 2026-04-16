package admin_user

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	UpdateUser(id int64, user *entity.UpdateUser) error
	UpdateUserRole(id int64, role string) error
}

type PassHash interface {
	EncryptPassword(password string) (string, error)
}

type Validate interface {
	ValidateAdmin(ctx context.Context, sessionID, accessToken string) (*entity.Payload, error)
}
