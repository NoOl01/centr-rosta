package admin

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository/transaction"
	"context"
)

const (
	timeLayout = "02-01-2006"
)

type UseCaseAdmin interface {
	TransactionStatsByTimePeriod(cxt context.Context, accessToken, sessionID, fromStr, toStr string) (*dto.TransactionStats, error)
}

type useCaseAdmin struct {
	rt      transaction.RepositoryTransaction
	session session.RepositorySession
}

func NewUseCaseAdmin(rt transaction.RepositoryTransaction, session session.RepositorySession) UseCaseAdmin {
	return &useCaseAdmin{
		rt:      rt,
		session: session,
	}
}
