package auth

import (
	"centr_rosta/internal/domain/usecase/auth"
)

type HandlerAuth struct {
	ua auth.UseCaseAuth
}

func NewHandlerAuth(ua auth.UseCaseAuth) *HandlerAuth {
	return &HandlerAuth{ua: ua}
}
