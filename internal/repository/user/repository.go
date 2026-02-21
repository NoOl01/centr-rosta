package user

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/repository/models"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	CreateUser(user *models.User) error
	UpdateUser(id int64, user dto.User) error
	UpdateUserRole(id int64, role string) error
	DeleteUser(id int64) error
	GetUseById(id int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type repositoryUser struct {
	Db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) RepositoryUser {
	return &repositoryUser{Db: db}
}
