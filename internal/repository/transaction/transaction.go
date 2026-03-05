package transaction

import (
	"centr_rosta/internal/repository/models"
	"time"
)

func (rt *repositoryTransaction) TransactionsByTimePeriod(from, to time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := rt.Db.Preload("User").Preload("Lesson").Where("created_at >= ? AND created_at <= ?", from, to).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
