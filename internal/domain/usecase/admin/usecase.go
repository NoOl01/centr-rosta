package admin

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

const (
	timeLayout = "02-01-2006"
)

type UseCaseAdmin interface {
	TransactionStatsByTimePeriod(ctx context.Context, accessToken, sessionID, fromStr, toStr string) (*[]entity.Transaction, float64, error)
}

type useCaseAdmin struct {
	rt       TransactionRepository
	validate Validate
}

func NewUseCaseAdmin(rt TransactionRepository, validate Validate) UseCaseAdmin {
	return &useCaseAdmin{
		rt:       rt,
		validate: validate,
	}
}
