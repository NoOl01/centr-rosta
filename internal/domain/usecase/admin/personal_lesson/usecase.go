package personal_lesson

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

type UseCaseAdminPersonalLesson interface {
	GetPersonalLessonsRequests(ctx context.Context, sessionID, accessToken string) ([]entity.PersonalLesson, error)
	ApprovePersonalLesson(ctx context.Context, sessionID, accessToken string, personalLessonID int64, time string) error
	CancelPersonalLesson(ctx context.Context, sessionID, accessToken string, personalLessonID int64) error
}

type useCaseAdminPersonalLesson struct {
	plr      PersonalLessonRepository
	validate Validate
}

func NewUseCaseAdminPersonalLesson(plr PersonalLessonRepository, validate Validate) UseCaseAdminPersonalLesson {
	return &useCaseAdminPersonalLesson{
		plr:      plr,
		validate: validate,
	}
}
