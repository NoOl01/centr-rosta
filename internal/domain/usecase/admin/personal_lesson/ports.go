package personal_lesson

import (
	"centr_rosta/internal/domain/entity"
	"context"
	"time"
)

type PersonalLessonRepository interface {
	Create(lessonID, userID int64, from, to time.Time) error
	Get() ([]entity.PersonalLesson, error)
	Update(personalLesson *entity.PersonalLesson) error
}

type Validate interface {
	ValidateAdmin(ctx context.Context, sessionID, accessToken string) (*entity.Payload, error)
}
