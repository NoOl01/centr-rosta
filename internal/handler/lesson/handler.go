package lesson

import "centr_rosta/internal/domain/usecase/lesson"

type HandlerLesson struct {
	ul lesson.UseCaseLesson
}

func NewHandlerLesson(ul lesson.UseCaseLesson) *HandlerLesson {
	return &HandlerLesson{
		ul: ul,
	}
}
