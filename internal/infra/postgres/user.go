package repository

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/internal/infra/postgres/models"
	"centr_rosta/pkg/logger"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetUsers() ([]entity.User, error) {
	var dbUsers []models.User
	if err := ur.db.Find(&dbUsers).Error; err != nil {
		return nil, errs.DbInternalError
	}

	var users []entity.User
	for _, u := range dbUsers {
		users = append(users, entity.User{
			ID:        &u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Role:      &u.Role,
		})
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(user *entity.User) error {
	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  *user.Password,
		CreatedAt: time.Now(),
	}

	if err := ur.db.Create(&newUser).Error; err != nil {
		logger.Log.Error(log_names.UserRepository, err.Error())
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errs.AlreadyExists
		}
		return errs.DbInternalError
	}

	user.ID = &newUser.ID
	user.Role = &newUser.Role

	return nil
}

func (ur *UserRepository) UpdateUser(id int64, user *entity.UpdateUser) error {
	newUser := models.User{
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
		Password:  *user.Password,
	}

	if err := ur.db.Where("id = ?", id).Updates(&newUser).Error; err != nil {
		logger.Log.Error(log_names.UserRepository, err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.RecordNotFound
		}
		return errs.DbInternalError
	}

	return nil
}

func (ur *UserRepository) UpdateUserRole(id int64, role string) error {
	if err := ur.db.Model(&models.User{}).Where("id = ?", id).Update("role", role).Error; err != nil {
		logger.Log.Error(log_names.UserRepository, err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.RecordNotFound
		}
		return errs.DbInternalError
	}

	return nil
}

func (ur *UserRepository) DeleteUser(id int64) error {
	if err := ur.db.Delete(&models.User{}, id).Error; err != nil {
		logger.Log.Error(log_names.UserRepository, err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.RecordNotFound
		}
		return errs.DbInternalError
	}

	return nil
}

func (ur *UserRepository) GetUserById(id int64) (*entity.User, error) {
	var dbUser models.User

	if err := ur.db.Where("id = ?", id).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RecordNotFound
		}
		return nil, errs.DbInternalError
	}

	user := entity.User{
		ID:        &dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Password:  &dbUser.Password,
		Role:      &dbUser.Role,
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var dbUser models.User

	if err := ur.db.Where("email = ?", email).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RecordNotFound
		}
		return nil, errs.DbInternalError
	}

	user := entity.User{
		ID:        &dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Password:  &dbUser.Password,
		Role:      &dbUser.Role,
	}

	return &user, nil
}
