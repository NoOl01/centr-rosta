package admin_user

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

func (uau *useCaseAdminUser) GetUsers(ctx context.Context, sessionID, accessToken string) ([]entity.User, error) {
	if err := uau.validate.ValidateAdmin(ctx, sessionID, accessToken); err != nil {
		return nil, err
	}

	users, err := uau.ur.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uau *useCaseAdminUser) ResetPassword(ctx context.Context, sessionID, accessToken string, userID int64) (*string, error) {
	if err := uau.validate.ValidateAdmin(ctx, sessionID, accessToken); err != nil {
		return nil, err
	}

	//todo

	return nil, nil
}

func (uau *useCaseAdminUser) UpdateRole(ctx context.Context, sessionID, accessToken, roleName string, userID int64) error {
	if err := uau.validate.ValidateAdmin(ctx, sessionID, accessToken); err != nil {
		return err
	}

	return uau.ur.UpdateUserRole(userID, roleName)
}
