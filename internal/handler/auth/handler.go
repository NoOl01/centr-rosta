package auth

import (
	"centr_rosta/internal/domain/usecase/auth"
)

type HandlerAuth struct {
	ua auth.IUseCaseAuth
}

func NewHandlerAuth(ua auth.IUseCaseAuth) *HandlerAuth {
	return &HandlerAuth{ua: ua}
}
