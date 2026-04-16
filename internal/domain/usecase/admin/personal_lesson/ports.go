package personal_lesson

import (
	"centr_rosta/internal/domain/entity"
	"time"
)

type PersonalLessonRepository interface {
	Create(lessonID, userID int64, from, to time.Time) error
	Get() ([]entity.PersonalLesson, error)
	Update(personalLesson *entity.PersonalLesson) error
}
