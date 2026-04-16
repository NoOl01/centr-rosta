package personal_lesson

import "centr_rosta/internal/domain/usecase/admin/personal_lesson"

type AdminPersonalLessonHandler struct {
	upcl personal_lesson.UseCaseAdminPersonalLesson
}

func NewAdminPersonalLessonHandler(ucpl personal_lesson.UseCaseAdminPersonalLesson) *AdminPersonalLessonHandler {
	return &AdminPersonalLessonHandler{
		upcl: ucpl,
	}
}
