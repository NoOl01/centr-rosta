package repository

import (
	"centr_rosta/internal/dto"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user dto.User) error
	UpdateUser(user dto.User) error
	UpdateUserRole(user dto.UpdateUserRole) error
	DeleteUser(id int64) error
}

type repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{Db: db}
}
