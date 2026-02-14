package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/utils/jwt"
	"centr_rosta/internal/utils/pass_hash"
	"context"
	"strconv"
)

func (s *serviceAuth) Login(ctx context.Context, user dto.Login) (string, string, string, error) {
	dbUser, err := s.repo.GetUser(user.Email)
	if err != nil {
		return "", "", "", err
	}

	if err := pass_hash.CheckPass(user.Password, dbUser.Password); err != nil {
		return "", "", "", err
	}

	userId := strconv.FormatInt(dbUser.ID, 10)

	payLoad := jwt.Payload{
		UserId: userId,
		Role:   dbUser.Role,
	}

	accessToken, refreshToken, err := jwt.GenerateToken(payLoad)
	if err != nil {
		return "", "", "", err
	}

	newSession := session.Session{
		UserID:       userId,
		DeviceToken:  "",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	sessionId, err := s.session.Create(ctx, newSession)
	if err != nil {
		return "", "", "", err
	}

	return accessToken, refreshToken, sessionId, nil
}
