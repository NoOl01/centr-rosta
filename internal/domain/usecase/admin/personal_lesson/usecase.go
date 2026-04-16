package personal_lesson

import (
	"centr_rosta/internal/domain/entity"
	"centr_rosta/internal/domain/usecase/validate"
	"context"
)

type UseCasePersonalLesson interface {
	GetPersonalLessonsRequests(ctx context.Context, sessionID, accessToken string) ([]entity.PersonalLesson, error)
	ApprovePersonalLesson(ctx context.Context, sessionID, accessToken string, personalLessonID int64, time string) error
	CancelPersonalLesson(ctx context.Context, sessionID, accessToken string, personalLessonID int64) error
}

type useCasePersonalLesson struct {
	plr      PersonalLessonRepository
	validate validate.Validate
}

func NewUseCasePersonalLesson(plr PersonalLessonRepository, validate validate.Validate) UseCasePersonalLesson {
	return &useCasePersonalLesson{
		plr:      plr,
		validate: validate,
	}
}
