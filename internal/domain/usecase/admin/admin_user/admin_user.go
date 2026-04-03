package admin_user

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/domain/entity"
	"context"
)

func (uau *useCaseAdminUser) GetUsers(ctx context.Context, sessionID, accessToken string) ([]entity.User, error) {
	session, err := uau.session.Get(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errs.SessionNotFound
	}

	if session.AccessToken != accessToken {
		return nil, errs.InvalidToken
	}

	payload, err := uau.jwt.ValidateJwt(accessToken)
	if err != nil {
		return nil, err
	}

	if payload.Role != entity.AdminRole {
		return nil, errs.AccessDenied
	}

	users, err := uau.ur.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
