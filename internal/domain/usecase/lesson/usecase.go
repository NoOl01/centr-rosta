package lesson

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type UseCaseLesson interface {
	GetLessons() ([]entity.Lesson, error)
	GetLessonByID(id int64) (*entity.Lesson, error)
	CreateLesson(ctx context.Context, sessionID, accessToken string, lesson *entity.Lesson) error
	UpdateLesson(ctx context.Context, sessionID, accessToken string, lesson *entity.Lesson) error
}

type useCaseLesson struct {
	rl       RepositoryLesson
	rs       RepositorySession
	validate Validate
}

func NewUseCaseLesson(rl RepositoryLesson, rs RepositorySession, validate Validate) UseCaseLesson {
	return &useCaseLesson{
		rl:       rl,
		rs:       rs,
		validate: validate,
	}
}
