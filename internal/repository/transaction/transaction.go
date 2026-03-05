package transaction

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/repository/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (rt *repositoryTransaction) TransactionsByTimePeriod(from, to time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := rt.Db.Preload("User").Preload("Lesson").Where("created_at >= ? AND created_at <= ?", from, to).Find(&transactions).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RecordNotFound
		}
		return nil, errs.DbInternalError
	}

	return transactions, nil
}
