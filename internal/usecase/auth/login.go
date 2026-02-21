package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/utils/pass_hash"
	"context"
)

func (ua *useCaseAuth) Login(ctx context.Context, user dto.Login) (string, string, string, error) {
	dbUser, err := ua.ru.GetUserByEmail(user.Email)
	if err != nil {
		return "", "", "", err
	}

	if err := pass_hash.CheckPass(user.Password, dbUser.Password); err != nil {
		return "", "", "", err
	}

	return ua.createSession(ctx, dbUser.ID, dbUser.Role)
}
