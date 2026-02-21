package user

import (
	"centr_rosta/internal/consts"
	"centr_rosta/internal/dto"
	"centr_rosta/internal/repository/models"
	"centr_rosta/pkg/logger"
)

func (ru *repositoryUser) CreateUser(user *models.User) error {
	if err := ru.Db.Create(user).Error; err != nil {
		logger.Log.Error(consts.UserRepository, err.Error())
		return err
	}

	return nil
}

func (ru *repositoryUser) UpdateUser(id int64, user dto.User) error {
	newUser := models.User{
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
		Password:  *user.Password,
	}

	if err := ru.Db.Where("id = ?", id).Updates(&newUser).Error; err != nil {
		logger.Log.Error(consts.UserRepository, err.Error())
		return err
	}

	return nil
}

func (ru *repositoryUser) UpdateUserRole(id int64, role string) error {
	if err := ru.Db.Model(&models.User{}).Where("id = ?", id).Update("role", role).Error; err != nil {
		logger.Log.Error(consts.UserRepository, err.Error())
		return err
	}

	return nil
}

func (ru *repositoryUser) DeleteUser(id int64) error {
	if err := ru.Db.Delete(&models.User{}, id).Error; err != nil {
		logger.Log.Error(consts.UserRepository, err.Error())
		return err
	}

	return nil
}

func (ru *repositoryUser) GetUserById(id int64) (*models.User, error) {
	var user models.User
	if err := ru.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ru *repositoryUser) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ru.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
