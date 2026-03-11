package repository

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/internal/infra/postgres/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (tr *TransactionRepository) TransactionsByTimePeriod(from, to time.Time) ([]entity.Transaction, error) {
	var dbTransactions []models.Transaction
	err := tr.db.Preload("User").Preload("LessonName").Where("created_at >= ? AND created_at <= ?", from, to).Find(&dbTransactions).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RecordNotFound
		}
		return nil, errs.DbInternalError
	}

	transactions := make([]entity.Transaction, 0, len(dbTransactions))
	for _, trn := range dbTransactions {
		transactions = append(transactions, entity.Transaction{
			UserID: trn.UserID,
			User: entity.User{
				FirstName: trn.User.FirstName,
				LastName:  trn.User.LastName,
				Email:     trn.User.Email,
			},
			Amount:   trn.Amount,
			Type:     trn.Type,
			LessonID: trn.LessonID,
			Lesson: entity.LessonName{
				Name: trn.Lesson.Name,
			},
		})
	}

	return transactions, nil
}
