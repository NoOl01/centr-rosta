package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/repository/models"

	"gorm.io/gorm"
)

type RepositoryAuth interface {
	CreateUser(user *models.User) error
	UpdateUser(id int64, user dto.User) error
	UpdateUserRole(id int64, role string) error
	DeleteUser(id int64) error
	GetUser(email string) (*models.User, error)
}

type repositoryAuth struct {
	Db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) RepositoryAuth {
	return &repositoryAuth{Db: db}
}
