package repository

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/internal/infra/postgres/models"
	"errors"

	"gorm.io/gorm"
)

type LessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) *LessonRepository {
	return &LessonRepository{
		db: db,
	}
}

func (lr *LessonRepository) Create(lesson *entity.Lesson) error {
	newLesson := models.Lesson{
		Name:        lesson.Name,
		Description: lesson.Description,
	}

	if err := lr.db.Create(&newLesson).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errs.AlreadyExists
		}
		return errs.DbInternalError
	}

	lesson.ID = &newLesson.ID
	return nil
}

func (lr *LessonRepository) GetAll() ([]*entity.Lesson, error) {
	var dbLessons []models.Lesson

	if err := lr.db.Find(&dbLessons).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RecordNotFound
		}
		return nil, errs.DbInternalError
	}

	lessons := make([]*entity.Lesson, 0, len(dbLessons))
	for _, lesson := range dbLessons {
		lessons = append(lessons, &entity.Lesson{
			ID:          &lesson.ID,
			Name:        lesson.Name,
			Description: lesson.Description,
		})
	}

	return lessons, nil
}

func (lr *LessonRepository) GetByID(id int64) (*entity.Lesson, error) {
	var dbLesson models.Lesson

	if err := lr.db.Where("id = ?", id).First(&dbLesson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RecordNotFound
		}
		return nil, errs.DbInternalError
	}

	lesson := &entity.Lesson{
		ID:          &dbLesson.ID,
		Name:        dbLesson.Name,
		Description: dbLesson.Description,
	}

	return lesson, nil
}
