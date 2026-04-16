package lesson

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type RepositoryLesson interface {
	Create(lesson *entity.Lesson) error
	Update(lessonID int64, lesson *entity.Lesson) error
	GetAll() ([]*entity.Lesson, error)
	GetByID(id int64) (*entity.Lesson, error)
}

type RepositorySession interface {
	Get(ctx context.Context, sessionID string) (*entity.Session, error)
}

type Validate interface {
	Validate(ctx context.Context, sessionID, accessToken string) (*entity.Payload, error)
	ValidateAdmin(ctx context.Context, sessionID, accessToken string) (*entity.Payload, error)
}
