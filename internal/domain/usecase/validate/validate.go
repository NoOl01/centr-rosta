package validate

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/domain/entity"
	"context"
)

type Validate interface {
	Validate(ctx context.Context, sessionID, accessToken string) (*entity.Payload, error)
	ValidateAdmin(ctx context.Context, sessionID, accessToken string) error
}

type validate struct {
	session SessionRepository
	jwt     Jwt
}

func NewValidate(session SessionRepository, jwt Jwt) Validate {
	return &validate{
		session: session,
		jwt:     jwt,
	}
}

func (v *validate) Validate(ctx context.Context, sessionID, accessToken string) (*entity.Payload, error) {
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

func (v *validate) ValidateAdmin(ctx context.Context, sessionID, accessToken string) error {
	payload, err := v.Validate(ctx, sessionID, accessToken)
	if err != nil {
		return err
	}

	if payload.Role != entity.AdminRole {
		return errs.AccessDenied
	}

	return nil
}
