package transaction

import (
	"centr_rosta/internal/repository/models"
	"time"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	TransactionsByTimePeriod(from, to time.Time) ([]models.Transaction, error)
}

type repositoryTransaction struct {
	Db *gorm.DB
}

func NewRepositoryTransaction(db *gorm.DB) RepositoryTransaction {
	return &repositoryTransaction{Db: db}
}
