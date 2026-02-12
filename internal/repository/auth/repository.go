package auth

import (
	"centr_rosta/internal/dto"

	"gorm.io/gorm"
)

type RepositoryAuth interface {
	CreateUser(user dto.User) error
	UpdateUser(id int64, user dto.User) error
	UpdateUserRole(id int64, role string) error
	DeleteUser(id int64) error
}

type repositoryAuth struct {
	Db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) RepositoryAuth {
	return &repositoryAuth{Db: db}
}
