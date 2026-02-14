package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository/models"
	"centr_rosta/internal/utils/jwt"
	"centr_rosta/internal/utils/pass_hash"
	"context"
	"strconv"
)

func (s *serviceAuth) Register(ctx context.Context, user dto.User) (string, string, string, error) {
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

	if err := s.repo.CreateUser(&newUser); err != nil {
		return "", "", "", err
	}

	userID := strconv.FormatInt(newUser.ID, 10)

	newPayload := jwt.Payload{
		UserId: userID,
		Role:   newUser.Role,
	}

	accessToken, refreshToken, err := jwt.GenerateToken(newPayload)
	if err != nil {
		return "", "", "", err
	}

	newSession := session.Session{
		UserID:      userID,
		DeviceToken: "",
	}

	sessionId, err := s.session.Create(ctx, newSession)
	if err != nil {
		return "", "", "", err
	}

	return accessToken, refreshToken, sessionId, nil
}
