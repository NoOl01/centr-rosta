package lesson

import (
	"centr_rosta/internal/domain/entity"
	"context"
)

func (ul *useCaseLesson) GetLessons() ([]entity.Lesson, error) {
	dbLessons, err := ul.rl.GetAll()
	if err != nil {
		return nil, err
	}

	lesson := make([]entity.Lesson, 0, len(dbLessons))
	for _, l := range dbLessons {
		lesson = append(lesson, entity.Lesson{
			ID:          l.ID,
			Name:        l.Name,
			Description: l.Description,
		})
	}

	return lesson, nil
}

func (ul *useCaseLesson) GetLessonByID(id int64) (*entity.Lesson, error) {
	dbLesson, err := ul.rl.GetByID(id)
	if err != nil {
		return nil, err
	}

	lesson := &entity.Lesson{
		ID:          dbLesson.ID,
		Name:        dbLesson.Name,
		Description: dbLesson.Description,
	}

	return lesson, nil
}

func (ul *useCaseLesson) CreateLesson(ctx context.Context, sessionID, accessToken string, lesson *entity.Lesson) error {
	if _, err := ul.validate.ValidateAdmin(ctx, sessionID, accessToken); err != nil {
		return err
	}

	return ul.rl.Create(lesson)
}

func (ul *useCaseLesson) UpdateLesson(ctx context.Context, sessionID, accessToken string, lesson *entity.Lesson) error {
	if _, err := ul.validate.ValidateAdmin(ctx, sessionID, accessToken); err != nil {
		return err
	}

	return ul.rl.Update(*lesson.ID, lesson)
}
