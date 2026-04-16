package repository

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/internal/infra/postgres/helper"
	"centr_rosta/internal/infra/postgres/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type PersonalLessonRepository struct {
	db *gorm.DB
}

func NewPersonalLessonRepository(db *gorm.DB) *PersonalLessonRepository {
	return &PersonalLessonRepository{
		db: db,
	}
}

func (plr *PersonalLessonRepository) Create(lessonID, userID int64, from, to time.Time) error {
	personalLesson := models.PersonalLesson{
		LessonID:          lessonID,
		UserID:            userID,
		EstimatedTimeFrom: from,
		EstimatedTimeTo:   to,
	}

	if err := plr.db.Create(&personalLesson).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errs.AlreadyExists
		}
		return errs.DbInternalError
	}

	return nil
}

func (plr *PersonalLessonRepository) Get() ([]entity.PersonalLesson, error) {
	var dbPersonalLessons []models.PersonalLesson
	var personalLessons []entity.PersonalLesson

	if err := plr.db.Preload("User").Preload("Lesson").Find(&dbPersonalLessons).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RecordNotFound
		}
		return nil, errs.DbInternalError
	}

	for _, dbLesson := range dbPersonalLessons {
		personalLessons = append(personalLessons, entity.PersonalLesson{
			ID:       &dbLesson.ID,
			LessonID: &dbLesson.LessonID,
			Lesson: &entity.Lesson{
				ID:          &dbLesson.Lesson.ID,
				Name:        dbLesson.Lesson.Name,
				Description: dbLesson.Lesson.Description,
			},
			UserID: &dbLesson.UserID,
			User: &entity.User{
				ID:        &dbLesson.User.ID,
				FirstName: dbLesson.User.FirstName,
				LastName:  dbLesson.User.LastName,
				Email:     dbLesson.User.Email,
			},
			EstimatedTimeFrom: &dbLesson.EstimatedTimeFrom,
			EstimatedTimeTo:   &dbLesson.EstimatedTimeTo,
			Status:            &dbLesson.Status,
		})
	}

	return personalLessons, nil
}

func (plr *PersonalLessonRepository) Update(personalLesson *entity.PersonalLesson) error {
	var dbPersonalLesson models.PersonalLesson

	if personalLesson.ID == nil {
		return errs.InvalidBody
	}

	if err := plr.db.Preload("User").Preload("Lesson").Where("id = ?", *personalLesson.ID).Find(&dbPersonalLesson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.RecordNotFound
		}
		return errs.DbInternalError
	}

	newPersonalLesson := helper.UpdatePersonalLessonStructBuilder(dbPersonalLesson, *personalLesson)

	if err := plr.db.Where("id = ?", newPersonalLesson.ID).Updates(&newPersonalLesson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.RecordNotFound
		}
		return errs.DbInternalError
	}

	return nil
}

func (plr *PersonalLessonRepository) Delete(personalLessonID int64) error {
	return plr.db.Where("id = ?", personalLessonID).Delete(&models.PersonalLesson{}).Error
}
