package admin_user

import (
	"centr_rosta/internal/domain/entity"
)

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	UpdateUserRole(id int64, role string) error
}
