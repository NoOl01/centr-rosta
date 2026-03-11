package admin

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

const (
	timeLayout = "02-01-2006"
)

type UseCaseAdmin interface {
	TransactionStatsByTimePeriod(cxt context.Context, accessToken, sessionID, fromStr, toStr string) (*[]entity.Transaction, float64, error)
}

type useCaseAdmin struct {
	rt      TransactionRepository
	session SessionRepository
	jwt     Jwt
}

func NewUseCaseAdmin(rt TransactionRepository, session SessionRepository, jwt Jwt) UseCaseAdmin {
	return &useCaseAdmin{
		rt:      rt,
		session: session,
		jwt:     jwt,
	}
}
