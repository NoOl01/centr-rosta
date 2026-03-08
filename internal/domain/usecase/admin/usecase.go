package admin

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

const (
	timeLayout = "02-01-2006"
)

type IUseCaseAdmin interface {
	TransactionStatsByTimePeriod(cxt context.Context, accessToken, sessionID, fromStr, toStr string) (*[]entity.Transaction, float64, error)
}

type useCaseAdmin struct {
	rt      ITransactionRepository
	session ISessionRepository
	jwt     IJwt
}

func NewUseCaseAdmin(rt ITransactionRepository, session ISessionRepository, jwt IJwt) IUseCaseAdmin {
	return &useCaseAdmin{
		rt:      rt,
		session: session,
		jwt:     jwt,
	}
}
