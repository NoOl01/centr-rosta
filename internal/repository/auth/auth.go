package auth

import (
	"centr_rosta/internal/consts"
	"centr_rosta/internal/dto"
	"centr_rosta/internal/repository/models"
	"centr_rosta/pkg/logger"
)

func (r *repositoryAuth) CreateUser(user *models.User) error {
	if err := r.Db.Create(user).Error; err != nil {
		logger.Log.Error(consts.AuthRepository, err.Error())
		return err
	}

	return nil
}

func (r *repositoryAuth) UpdateUser(id int64, user dto.User) error {
	newUser := models.User{
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
		Password:  *user.Password,
	}

	if err := r.Db.Where("id = ?", id).Updates(&newUser).Error; err != nil {
		logger.Log.Error(consts.AuthRepository, err.Error())
		return err
	}

	return nil
}

func (r *repositoryAuth) UpdateUserRole(id int64, role string) error {
	if err := r.Db.Model(&models.User{}).Where("id = ?", id).Update("role", role).Error; err != nil {
		logger.Log.Error(consts.AuthRepository, err.Error())
		return err
	}

	return nil
}

func (r *repositoryAuth) DeleteUser(id int64) error {
	if err := r.Db.Delete(&models.User{}, id).Error; err != nil {
		logger.Log.Error(consts.AuthRepository, err.Error())
		return err
	}

	return nil
}

func (r *repositoryAuth) GetUser(email string) (*models.User, error) {
	var user models.User
	if err := r.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
