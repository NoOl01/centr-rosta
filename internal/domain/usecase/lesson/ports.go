package lesson

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type RepositoryLesson interface {
	Create(lesson *entity.Lesson) error
	GetAll() ([]*entity.Lesson, error)
	GetByID(id int64) (*entity.Lesson, error)
}

type RepositorySession interface {
	Get(ctx context.Context, sessionID string) (*entity.Session, error)
}
