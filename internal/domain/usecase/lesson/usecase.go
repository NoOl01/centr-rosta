package lesson

import "centr_rosta/internal/domain/entity"

type UseCaseLesson interface {
	GetLessons() ([]entity.Lesson, error)
	GetLessonByID(id int64) (*entity.Lesson, error)
}

type useCaseLesson struct {
	rl RepositoryLesson
	rs RepositorySession
}

func NewUseCaseLesson(rl RepositoryLesson, rs RepositorySession) UseCaseLesson {
	return &useCaseLesson{
		rl: rl,
		rs: rs,
	}
}
