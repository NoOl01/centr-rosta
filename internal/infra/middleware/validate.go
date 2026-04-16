package middleware

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/internal/infra/jwt"
	"centr_rosta/internal/infra/redis"
	"context"
)

type ValidateMiddleWare struct {
	session redis.SessionRepository
	jwt     jwt.ServiceJwt
}

func NewValidateMiddleWare(session redis.SessionRepository, jwt jwt.ServiceJwt) *ValidateMiddleWare {
	return &ValidateMiddleWare{
		session: session,
		jwt:     jwt,
	}
}

func (v *ValidateMiddleWare) Validate(ctx context.Context, sessionID, accessToken string) (*entity.Payload, error) {
	session, err := v.session.Get(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errs.SessionNotFound
	}

	if session.AccessToken != accessToken {
		return nil, errs.InvalidToken
	}

	payload, err := v.jwt.ValidateJwt(accessToken)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (v *ValidateMiddleWare) ValidateAdmin(ctx context.Context, sessionID, accessToken string) (*entity.Payload, error) {
	payload, err := v.Validate(ctx, sessionID, accessToken)
	if err != nil {
		return nil, err
	}

	if payload.Role != entity.AdminRole {
		return nil, errs.AccessDenied
	}

	return payload, nil
}
