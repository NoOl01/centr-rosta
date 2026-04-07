package admin_user

import (
	"centr_rosta/internal/domain/entity"
)

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	UpdateUser(id int64, user *entity.UpdateUser) error
	UpdateUserRole(id int64, role string) error
}

type PassHash interface {
	EncryptPassword(password string) (string, error)
}
