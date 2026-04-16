package personal_lesson

import "centr_rosta/internal/domain/usecase/admin/personal_lesson"

type AdminPersonalLessonHandler struct {
	upcl personal_lesson.UseCasePersonalLesson
}

func NewAdminPersonalLessonHandler(ucpl personal_lesson.UseCasePersonalLesson) *AdminPersonalLessonHandler {
	return &AdminPersonalLessonHandler{
		upcl: ucpl,
	}
}
