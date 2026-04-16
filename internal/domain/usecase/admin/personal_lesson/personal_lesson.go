package personal_lesson

import (
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/pkg/logger"
	"context"
	"time"
)

const layout = "02.01.2006 15:04:05"

func (ucpl *useCasePersonalLesson) GetPersonalLessonsRequests(ctx context.Context, sessionID, accessToken string) ([]entity.PersonalLesson, error) {
	_, err := ucpl.validate.ValidateAdmin(ctx, sessionID, accessToken)
	if err != nil {
		return nil, err
	}

	return ucpl.plr.Get()
}

func (ucpl *useCasePersonalLesson) ApprovePersonalLesson(ctx context.Context, sessionID, accessToken string, personalLessonID int64, newTime string) error {
	_, err := ucpl.validate.ValidateAdmin(ctx, sessionID, accessToken)
	if err != nil {
		return err
	}

	exactTime, err := time.Parse(layout, newTime)
	if err != nil {
		logger.Log.Error(log_names.ApproveUseCase, err.Error())
		return err
	}

	personalLesson := entity.PersonalLesson{
		ID:        &personalLessonID,
		ExactTime: &exactTime,
		Status:    new(entity.Approved),
	}

	return ucpl.plr.Update(&personalLesson)
}

func (ucpl *useCasePersonalLesson) CancelPersonalLesson(ctx context.Context, sessionID, accessToken string, personalLessonID int64) error {
	_, err := ucpl.validate.ValidateAdmin(ctx, sessionID, accessToken)
	if err != nil {
		return err
	}

	return ucpl.CancelPersonalLesson(ctx, sessionID, accessToken, personalLessonID)
}
