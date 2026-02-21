package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/repository/models"
	"centr_rosta/internal/utils/pass_hash"
	"context"
)

func (ua *useCaseAuth) Register(ctx context.Context, user dto.User) (string, string, string, error) {
	var err error
	*user.Password, err = pass_hash.EncryptPassword(*user.Password)
	if err != nil {
		return "", "", "", err
	}

	newUser := models.User{
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
		Password:  *user.Password,
	}

	if err := ua.ru.CreateUser(&newUser); err != nil {
		return "", "", "", err
	}

	return ua.createSession(ctx, newUser.ID, newUser.Role)
}
